package mysql

import (
	"context"
	"net/http/httptrace"
)

func NewTrace(ctx context.Context) context.Context {
	return httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		DNSStart:     dnsStart(ctx),
		DNSDone:      dnsDone(ctx),
		ConnectStart: connectStartTrace(ctx),
		ConnectDone:  connectDone(ctx),
	})
}

func dnsStart(ctx context.Context) func(httptrace.DNSStartInfo) {
	return func(info httptrace.DNSStartInfo) {
		infoLog.Print(ctx, "DNS start with info: %v", info)
	}
}

func dnsDone(ctx context.Context) func(httptrace.DNSDoneInfo) {
	return func(info httptrace.DNSDoneInfo) {
		infoLog.Print(ctx, "DNS done with info: %v", info)
	}
}

func connectStartTrace(ctx context.Context) func(network string, addr string) {
	return func(network string, addr string) {
		infoLog.Print(ctx, "Connect start at network %q to address %q", network, addr)
	}
}

func connectDone(ctx context.Context) func(network string, addr string, err error) {
	return func(network string, addr string, err error) {
		if err != nil {
			errLog.Print(ctx, "Connect done at network %q to address %q with err: %s", network, addr, err)
			return
		}
		infoLog.Print(ctx, "Connect done at network %q to address %q", network, addr)
	}
}
