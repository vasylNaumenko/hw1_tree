/*
 * Copyright (c) 2021. Vasyl Naumenko
 */

package main

import (
	"fmt"
	"strconv"

	"git.ooo.ua/vipcoin/go-common/cipher"
)

func main() {
	cipherKey := []byte("LKHlhb899Y09olUi")
	uids := []int64{1, 2}

	for _, uid := range uids {
		// Getting user hash from user id number
		userHash, _ := cipher.Encrypt(cipherKey, strconv.FormatInt(uid, 10))
		fmt.Printf("user id: %v hash: %s\n", uid, userHash)
	}
}
