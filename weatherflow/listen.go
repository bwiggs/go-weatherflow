package weatherflow

import (
	"net"
	"strings"

	"golang.org/x/exp/slog"
)

const BuffSize = 512

func Listen() {
	addr := &net.UDPAddr{Port: 50222}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		b := make([]byte, BuffSize)

		n, a, err := conn.ReadFrom(b)
		if err != nil {
			slog.Error("Error reading packet: %s", err)
		}

		dat := strings.TrimSpace(string(b))

		_ = n
		_ = a
		slog.Debug("packet received")
		slog.Info(dat)
	}
}
