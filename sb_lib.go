// v0.03

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

func main() {
	fmt.Println("Roku SoundBridge GFX Library")

	// Connect to 192.168.0.228 port 4444
	// Send 'sketch\r\n'
	// Connected!
	conn, err := net.Dial("tcp", "192.168.0.228:4444")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Connected!")

	// Enter 'sketch' mode
	time.Sleep(1 * time.Second) // heh
	data := []byte("sketch\r\n")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	
	// Draw
	fmt.Println("Sending sketch")
	data = []byte("text 5 5 \"hello\"\r\n")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1 * time.Second)

	// Exit
	fmt.Println("Exiting...")
	data = []byte("exit\r\nexit\r\n")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
}

/*
// Connectivity
func sb_connect(conn net.Conn) {
	// Connect to 192.168.0.228 port 4444
	// Send 'sketch\r\n'
	// Connected!
	conn, err := net.Dial("tcp", "192.168.0.228:4444")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Connected!")
	// defer conn.Close()
}

func sb_disconnect() {}
func sb_sync() {} // Keep-Alive?

// Graphics Functions
func point(int x, int y) {}
func line() {}
func square(int x_1, int y_1, int x_2, int y_2, bool filled){}
func interleaving_lines() {
	// | ||| | ||| | ||| this kinda stuff. lines flip back and forth between on and off
}

// Demo Games (not playable.... yet)
func pong(){}
func snake(){}
*/
