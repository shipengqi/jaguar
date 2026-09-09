package cache

import "strconv"

func getaddrs(opts *Options) (addrs []string) {
	if len(opts.Addrs) > 0 {
		addrs = opts.Addrs
	}

	if len(addrs) == 0 && opts.Port != 0 {
		addr := opts.Host + ":" + strconv.Itoa(opts.Port)
		addrs = append(addrs, addr)
	}

	return addrs
}
