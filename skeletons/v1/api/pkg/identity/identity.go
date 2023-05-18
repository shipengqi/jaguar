package identity

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	// ErrMissingSecretKey indicates Secret key is required
	ErrMissingSecretKey = errors.New("secret key is required")

	// ErrForbidden when HTTP status 403 is given
	ErrForbidden = errors.New("you don't have permission to access this resource")

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = errors.New("token is expired")

	// ErrEmptyAuthHeader can be thrown if authing with an HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("auth header is invalid")

	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = errors.New("incorrect Username or Password")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = errors.New("query token is empty")

	// ErrEmptyCookieToken can be thrown if authing with a cookie, the token cookie is empty
	ErrEmptyCookieToken = errors.New("cookie token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = errors.New("parameter token is empty")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")

	// ErrNoPrivKeyFile indicates that the given private key is unreadable
	ErrNoPrivKeyFile = errors.New("private key file unreadable")

	// ErrNoPubKeyFile indicates that the given public key is unreadable
	ErrNoPubKeyFile = errors.New("public key file unreadable")

	// ErrInvalidPrivKey indicates that the given private key is invalid
	ErrInvalidPrivKey = errors.New("private key invalid")

	// ErrInvalidPubKey indicates the the given public key is invalid
	ErrInvalidPubKey = errors.New("public key invalid")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = errors.New("missing exp field")

	// ErrWrongFormatOfExp field must be float64 format
	ErrWrongFormatOfExp = errors.New("exp must be float64 format")

	// ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingAuthenticatorFunc = errors.New("ginJWTMiddleware.Authenticator func is undefined")

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = errors.New("failed to create JWT Token")
)

const (
	DefaultIdentityKey = "identity"
)

// MapClaims type that uses the map[string]interface{} for JSON decoding
// This is the default claims type if you don't supply one
type MapClaims map[string]interface{}

// Middleware provides a Json-Web-Token authentication implementation.
type Middleware struct {
	// Realm name to display to the user. Required.
	Realm string

	// signing algorithm - possible values are HS256, HS384, HS512, RS256, RS384 or RS512
	// Optional, default is HS256.
	SigningAlgorithm string

	// Secret key used for signing. Required.
	Key []byte

	// Duration that a jwt token is valid. Optional, defaults to one hour.
	Timeout time.Duration

	// This field allows clients to refresh their token until MaxRefresh has passed.
	// Note that clients can refresh their token in the last moment of MaxRefresh.
	// This means that the maximum validity timespan for a token is TokenTime + MaxRefresh.
	// Optional, defaults to 0 meaning not refreshable.
	MaxRefresh time.Duration

	// Callback function that should perform the authentication of the user based on login info.
	// Must return user data as user identifier, it will be stored in Claim Array. Required.
	// Check error (e) to determine the appropriate error message.
	Authenticator func(c *gin.Context) (interface{}, error)

	// Callback function that should perform the authorization of the authenticated user. Called
	// only after an authentication success. Must return true on success, false on failure.
	// Optional, default to success.
	Authorizer func(data interface{}, c *gin.Context) bool

	// Callback function that will be called during login.
	// Using this function it is possible to add additional payload data to the jwt token.
	// The data is then made available during requests via c.Get("JWT_PAYLOAD").
	// Note that the payload is not encrypted.
	// The attributes mentioned on jwt.io can't be used as keys for the map.
	// Optional, by default no additional data will be set.
	PayloadFunc func(data interface{}) MapClaims

	// User can define own Unauthorized func.
	// code: HTTP status code
	Unauthorized func(ctx *gin.Context, code int, message string)

	// User can define own LoginResponse func.
	LoginResponse func(ctx *gin.Context, token string, expire time.Time)

	// User can define own LogoutResponse func.
	LogoutResponse func(ctx *gin.Context)

	// User can define own RefreshResponse func.
	RefreshResponse func(ctx *gin.Context, token string, expire time.Time)

	// Set the identity handler function
	IdentityHandler func(*gin.Context) interface{}

	// Set the identity key
	IdentityKey string

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// - "param:<name>"
	TokenLookup string

	// Optionally return the token as a cookie
	SendCookie bool

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName string

	// Duration that a cookie is valid. Optional, by default equals to Timeout value.
	CookieMaxAge time.Duration

	// Allow insecure cookies for development over http
	SecureCookie bool

	// Allow cookies to be accessed client side for development
	CookieHTTPOnly bool

	// Allow cookie domain change for development
	CookieDomain string

	// CookieName allow cookie name change for development
	CookieName string

	// CookieSameSite allow use http.SameSite cookie param
	// See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.
	CookieSameSite http.SameSite

	// Private key file for asymmetric algorithms
	PrivKeyFile string

	// Public key file for asymmetric algorithms
	PubKeyFile string

	// Private key
	privKey *rsa.PrivateKey

	// Public key
	pubKey *rsa.PublicKey
}

// New for check error with GinJWTMiddleware
func New(m *Middleware) (*Middleware, error) {
	if err := m.MiddlewareInit(); err != nil {
		return nil, err
	}

	return m, nil
}

// MiddlewareInit initialize jwt configs.
func (m *Middleware) MiddlewareInit() error {

	if m.TokenLookup == "" {
		m.TokenLookup = "header:Authorization"
	}

	if m.SigningAlgorithm == "" {
		m.SigningAlgorithm = "HS256"
	}

	if m.Timeout == 0 {
		m.Timeout = time.Hour
	}

	m.TokenHeadName = strings.TrimSpace(m.TokenHeadName)
	if len(m.TokenHeadName) == 0 {
		m.TokenHeadName = "Bearer"
	}

	if m.Authorizer == nil {
		m.Authorizer = func(data interface{}, c *gin.Context) bool {
			return true
		}
	}

	if m.Unauthorized == nil {
		m.Unauthorized = func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		}
	}

	if m.LoginResponse == nil {
		m.LoginResponse = func(c *gin.Context, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		}
	}

	if m.LogoutResponse == nil {
		m.LogoutResponse = func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		}
	}

	if m.RefreshResponse == nil {
		m.RefreshResponse = func(c *gin.Context, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		}
	}

	if m.IdentityKey == "" {
		m.IdentityKey = DefaultIdentityKey
	}

	if m.IdentityHandler == nil {
		m.IdentityHandler = func(c *gin.Context) interface{} {
			claims := ExtractClaims(c)
			return claims[m.IdentityKey]
		}
	}

	if m.Realm == "" {
		m.Realm = "identity jwt"
	}

	if m.CookieMaxAge == 0 {
		m.CookieMaxAge = m.Timeout
	}

	if m.CookieName == "" {
		m.CookieName = "X-AUTH-TOKEN"
	}

	if m.IsPublicKeyAlgo() {
		return m.readKeys()
	}

	if m.Key == nil {
		return ErrMissingSecretKey
	}
	return nil
}

// MiddlewareFunc makes Middleware implement the gin.HandlerFunc.
func (m *Middleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.GinMiddlewareImpl(c)
	}
}

// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
func (m *Middleware) LoginHandler(c *gin.Context) {
	if m.Authenticator == nil {
		m.unauthorized(c, http.StatusInternalServerError, ErrMissingAuthenticatorFunc)
		return
	}

	data, err := m.Authenticator(c)
	if err != nil {
		m.unauthorized(c, http.StatusUnauthorized, err)
		return
	}

	// Create the token
	claims := make(jwt.MapClaims)
	if m.PayloadFunc != nil {
		for key, value := range m.PayloadFunc(data) {
			claims[key] = value
		}
	}

	tokenString, expire, err := m.GenerateToken(claims)
	if err != nil {
		m.unauthorized(c, http.StatusUnauthorized, ErrFailedTokenCreation)
		return
	}

	// set cookie
	m.SetCookie(c, tokenString)

	m.LoginResponse(c, tokenString, expire)
}

// LogoutHandler can be used by clients to remove the jwt cookie (if set)
func (m *Middleware) LogoutHandler(c *gin.Context) {
	// clean auth cookie
	m.CleanCookie(c)

	m.LogoutResponse(c)
}

// RefreshHandler can be used to refresh a token. The token still needs to be valid on refresh.
// Shall be put under an endpoint that is using the GinJWTMiddleware.
// Reply will be of the form {"token": "TOKEN"}.
func (m *Middleware) RefreshHandler(c *gin.Context) {
	tokenString, expire, err := m.RefreshToken(c)
	if err != nil {
		m.unauthorized(c, http.StatusUnauthorized, err)
		return
	}

	m.RefreshResponse(c, tokenString, expire)
}

// RefreshToken refresh token and check if token is expired
func (m *Middleware) RefreshToken(c *gin.Context) (string, time.Time, error) {
	claims, err := m.CheckIfTokenExpire(c)
	if err != nil {
		return "", time.Now(), err
	}

	// Create the token
	tokenString, expire, err := m.GenerateToken(claims)
	if err != nil {
		return "", time.Now(), err
	}

	// set cookie
	m.SetCookie(c, tokenString)

	return tokenString, expire, nil
}

func (m *Middleware) GinMiddlewareImpl(c *gin.Context) {
	claims, err := m.GetClaimsFromJWT(c)
	if err != nil {
		m.unauthorized(c, http.StatusUnauthorized, err)
		return
	}

	if claims["exp"] == nil {
		m.unauthorized(c, http.StatusBadRequest, ErrMissingExpField)
		return
	}

	if _, ok := claims["exp"].(float64); !ok {
		m.unauthorized(c, http.StatusBadRequest, ErrWrongFormatOfExp)
		return
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		m.unauthorized(c, http.StatusUnauthorized, ErrExpiredToken)
		return
	}

	c.Set("JWT_PAYLOAD", claims)
	identity := m.IdentityHandler(c)

	if identity != nil {
		c.Set(m.IdentityKey, identity)
	}

	if !m.Authorizer(identity, c) {
		m.unauthorized(c, http.StatusForbidden, ErrForbidden)
		return
	}

	c.Next()
}

// GetClaimsFromJWT get claims from JWT token
func (m *Middleware) GetClaimsFromJWT(c *gin.Context) (MapClaims, error) {
	token, err := m.ParseToken(c)

	if err != nil {
		return nil, err
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims, nil
}

// ParseToken parse jwt token from gin context
func (m *Middleware) ParseToken(c *gin.Context) (*jwt.Token, error) {
	var token string
	var err error

	methods := strings.Split(m.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = m.jwtFromHeader(c, v)
		case "query":
			token, err = m.jwtFromQuery(c, v)
		case "cookie":
			token, err = m.jwtFromCookie(c, v)
		case "param":
			token, err = m.jwtFromParam(c, v)
		}
	}

	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(m.SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		if m.IsPublicKeyAlgo() {
			return m.pubKey, nil
		}

		// save token string if valid
		c.Set("JWT_TOKEN", token)

		return m.Key, nil
	})
}

// ParseTokenString parse jwt token string
func (m *Middleware) ParseTokenString(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(m.SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		if m.IsPublicKeyAlgo() {
			return m.pubKey, nil
		}

		return m.Key, nil
	})
}

// CheckIfTokenExpire check if token expire
func (m *Middleware) CheckIfTokenExpire(c *gin.Context) (jwt.MapClaims, error) {
	token, err := m.ParseToken(c)

	if err != nil {
		// If we receive an error, and the error is anything other than a single
		// ValidationErrorExpired, we want to return the error.
		// If the error is just ValidationErrorExpired, we want to continue, as we can still
		// refresh the token if it's within the MaxRefresh time.
		// (see https://github.com/appleboy/gin-jwt/issues/176)
		validationErr, ok := err.(*jwt.ValidationError)
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return nil, err
		}
	}

	claims := token.Claims.(jwt.MapClaims)

	origIat := int64(claims["orig_iat"].(float64))

	if origIat < time.Now().Add(-m.MaxRefresh).Unix() {
		return nil, ErrExpiredToken
	}

	return claims, nil
}

func (m *Middleware) SignedString(token *jwt.Token) (string, error) {
	var tokenString string
	var err error
	if m.IsPublicKeyAlgo() {
		tokenString, err = token.SignedString(m.privKey)
	} else {
		tokenString, err = token.SignedString(m.Key)
	}
	return tokenString, err
}

// ExtractClaims help to extract the JWT claims from gin.Context
func ExtractClaims(c *gin.Context) MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(MapClaims)
	}

	return claims.(MapClaims)
}

// ExtractClaimsFromToken help to extract the JWT claims from token
func ExtractClaimsFromToken(token *jwt.Token) MapClaims {
	if token == nil {
		return make(MapClaims)
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims
}

// GenerateToken method that clients can use to generate a jwt token.
func (m *Middleware) GenerateToken(data jwt.MapClaims) (string, time.Time, error) {
	token := jwt.New(jwt.GetSigningMethod(m.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range data {
		claims[key] = value
	}

	now := time.Now()
	expire := time.Now().UTC().Add(m.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = now.Unix()
	tokenString, err := m.SignedString(token)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expire, nil
}

func (m *Middleware) SetCookie(c *gin.Context, token string) {
	// set cookie
	if m.SendCookie {
		now := time.Now()
		expireCookie := now.Add(m.CookieMaxAge)
		maxage := int(expireCookie.Unix() - now.Unix())

		if m.CookieSameSite != 0 {
			c.SetSameSite(m.CookieSameSite)
		}

		c.SetCookie(
			m.CookieName,
			token,
			maxage,
			"/",
			m.CookieDomain,
			m.SecureCookie,
			m.CookieHTTPOnly,
		)
	}
}

func (m *Middleware) CleanCookie(c *gin.Context) {
	// clean auth cookie
	if m.SendCookie {
		if m.CookieSameSite != 0 {
			c.SetSameSite(m.CookieSameSite)
		}

		c.SetCookie(
			m.CookieName,
			"",
			-1,
			"/",
			m.CookieDomain,
			m.SecureCookie,
			m.CookieHTTPOnly,
		)
	}
}

func (m *Middleware) unauthorized(c *gin.Context, code int, err error) {
	c.Header("WWW-Authenticate", "JWT realm="+m.Realm)
	c.Abort()

	m.Unauthorized(c, code, err.Error())
}

func (m *Middleware) jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == m.TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func (m *Middleware) jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

func (m *Middleware) jwtFromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Cookie(key)

	if cookie == "" {
		return "", ErrEmptyCookieToken
	}

	return cookie, nil
}

func (m *Middleware) jwtFromParam(c *gin.Context, key string) (string, error) {
	token := c.Param(key)

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}

func (m *Middleware) readKeys() error {
	err := m.privateKey()
	if err != nil {
		return err
	}
	err = m.publicKey()
	if err != nil {
		return err
	}
	return nil
}

func (m *Middleware) privateKey() error {
	keyData, err := ioutil.ReadFile(m.PrivKeyFile)
	if err != nil {
		return ErrNoPrivKeyFile
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return ErrInvalidPrivKey
	}
	m.privKey = key
	return nil
}

func (m *Middleware) publicKey() error {
	keyData, err := ioutil.ReadFile(m.PubKeyFile)
	if err != nil {
		return ErrNoPubKeyFile
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return ErrInvalidPubKey
	}
	m.pubKey = key
	return nil
}

func (m *Middleware) IsPublicKeyAlgo() bool {
	switch m.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}
