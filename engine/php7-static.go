// Copyright 2016 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package engine

// #cgo CFLAGS: -I/usr/include/php/20151012 -Iinclude/php7 -Isrc/php7 -Iinclude
// #cgo CFLAGS: -I/usr/include/php/20151012 -I/usr/include/php/20151012/main -I/usr/include/php/20151012/TSRM -I/usr/include/php/20151012/Zend 
// #cgo CFLAGS: -I/usr/include/php/20151012/ext -I/usr/include/php/20151012/ext/date/lib
// #cgo LDFLAGS: -L/usr/lib/php/20151012  -L/usr/lib/x86_64-linux-gnu
// #cgo LDFLAGS: -lphp7.0 -lm -ldl -lcrypt -lresolv -lcrypt -lz -lpcre -lrt -lm -ldl -lnsl -lcurl -lxml2 -lssl -lcrypto -lcrypt -lcrypt -lmcrypt -lopcache
import "C"
