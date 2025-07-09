// v0.05
// Time for GitHub! I think under STAN PERSONAL???? like with BMC Virtual Disk stuff

// VFD Display:
/*
       x_=_0_______________________________x_=_280
y = 0  |                                         |
       |                                         |
y = 32 |_________________________________________|

*/

package main

import (
	"fmt"
	"net"
	"time"
)

var sb_conn net.Dialer

func main() {
	fmt.Println("Roku SoundBridge GFX Library")

	sb_conn := sb_connect() // This is the "net.Conn" part AFTER the "net.Dial/Dialer" is set up!! (likew "new_conn" in sb_connect()
	// Enter 'sketch' mode
	fmt.Fprintf(sb_conn, "sketch\r\n")
	
	// Draw
	fmt.Println("Sending sketch")
	fmt.Fprintf(sb_conn, "text 5 5 \"hello\"\r\n")
	time.Sleep(1 * time.Second)

	// Exit
	fmt.Println("Exiting...")
	fmt.Fprintf(sb_conn, "exit\r\nexit\r\n")

	defer sb_conn.Close()
}

// Connectivity
func sb_connect() (butts net.Conn){
	// Connect to 192.168.0.228 port 4444
	// Send 'sketch\r\n'
	// Connected!
	new_conn, err := sb_conn.Dial("tcp", "192.168.0.228:4444")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	time.Sleep(1 * time.Second) // heh
	fmt.Fprintf(new_conn, "sketch\r\n")
	fmt.Println("Connected!")

	// Return the net.Conn "handler"(?)
	return new_conn
}
