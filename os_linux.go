//go:build linux
// +build linux

package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/internal/wglinux"
	"golang.zx2c4.com/wireguard/wgctrl/internal/wguser"
)

// newClients configures wginternal.Clients for Linux systems.
func newClients() ([]wginternal.Client, error) {
	var clients []wginternal.Client

	// Linux has an in-kernel WireGuard implementation. Determine if it is
	// available and make use of it if so.
	kc, ok, err := wglinux.New()

	// The reason of my fork - if we've got an error when trying to use in-kernel implementation,
	// we should try to use userspace instead
	if err == nil && ok {
		clients = append(clients, kc)
	}

	// Although it isn't recommended to use userspace implementations on Linux,
	// it can be used. We make use of it in integration tests as well.
	uc, err := wguser.New()
	if err != nil {
		return nil, err
	}

	// Kernel devices seem to appear first in wg(8).
	clients = append(clients, uc)
	return clients, nil
}
