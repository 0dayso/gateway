package proxy

import (
	"github.com/fagongzi/gateway/conf"
	"net"
	"strings"
)

// record the http access log
// log format: $remoteip "$method $path HTTP/$proto" $code "$agent"

type XForwardForFilter struct {
	baseFilter
	config *conf.Conf
	proxy  *Proxy
}

func newXForwardForFilter(config *conf.Conf, proxy *Proxy) Filter {
	return XForwardForFilter{
		config: config,
		proxy:  proxy,
	}
}

func (self XForwardForFilter) Name() string {
	return FILTER_XFORWARD
}

func (self XForwardForFilter) Pre(c *filterContext) (statusCode int, err error) {
	if clientIP, _, err := net.SplitHostPort(c.req.RemoteAddr); err == nil {
		// If we aren't the first proxy retain prior
		// X-Forwarded-For information as a comma+space
		// separated list and fold multiple headers into one.
		if prior, ok := c.outreq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		c.outreq.Header.Set("X-Forwarded-For", clientIP)
	}

	return self.baseFilter.Pre(c)
}
