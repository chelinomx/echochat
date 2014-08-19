package main

import (
	"bytes"
	"net"
	"strings"
)

func strcat(s1 string, s2 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}

func GetUserByNick(nick string) *User {
	nick = strings.ToLower(nick)
	for _, k := range userlist {
		if strings.ToLower(k.nick) == nick {
			return k
		}
	}
	return nil
}

func GetIpFromConn(conn net.Conn) string {
	ip := conn.RemoteAddr().String()
	if !strings.HasPrefix(ip, "[") {
		//IPV4
		return strings.Split(ip, ":")[0]
	} else {
		//IPV6
		ip = strings.Split(ip, "]")[0]
		ip = strings.TrimPrefix(ip, "[")
		return ip
	}
}

func GetChannelByName(name string) *Channel {
	return chanlist[strings.ToLower(name)]
}

func SendToMany(msg string, list []*User) {
	users := make(map[*User]int)
	for _, j := range list {
		users[j] = 0
	}
	for j, _ := range users {
		j.SendLine(msg)
	}
}

func ValidChanName(name string) bool {
	if ChanHasBadChars(name) {
		return false
	}
	for _, k := range valid_chan_prefix {
		if strings.HasPrefix(name, k) {
			return true
		}
	}
	return false
}

//IMPORTANT: args must ABSOLUTELY be a valid privmsg command, or this will not work
//validity does not depend on leading ":", I don't care that much
func FormatMessageArgs(args []string) string {
	msg := strings.Join(args[2:], " ")
	msg = strings.TrimPrefix(msg, ":")
	return msg
}

func NickHasBadChars(nick string) bool {
	for _, k := range global_bad_chars {
		if strings.Contains(nick, k) {
			return true
		}
	}
	for _, k := range valid_chan_prefix {
		if strings.Contains(nick, k) {
			return true
		}
	}
	return false
}

func ChanHasBadChars(nick string) bool {
	for _, k := range global_bad_chars {
		if strings.Contains(nick, k) {
			return true
		}
	}
	return false
}

func ChanUserNone(name string) int {
	if GetChannelByName(name) != nil {
		return 1
	} else if GetUserByNick(name) != nil {
		return 2
	} else {
		return 0
	}
}

func WildcardMatch(text string, pattern string) bool {
	cards := strings.Split(pattern, "*")
	for _, card := range cards {
		index := strings.Index(text, card)
		if index == -1 {
			return false
		}
		text = text[index+len(card):]
	}
	return true
}
