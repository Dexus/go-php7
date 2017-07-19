// Copyright 2016 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package engine

// #cgo CFLAGS: -Iinclude/php7 -Isrc/php7 -Iinclude
// #cgo CFLAGS: -I/opt/php7/include/php -I/opt/php7/include/php/Zend -I/opt/php7/include/php/TSRM -I/opt/php7/include/php/main
// #cgo LDFLAGS: -L/opt/php7/lib -L/opt/curl/lib -L/opt/libmcrypt/lib -L/opt/zlib/lib -L/opt/openssl/lib -L/opt/libxml2/lib
// #cgo LDFLAGS: -L/opt/php7/lib/php/extensions/debug-non-zts-20160303
// #cgo LDFLAGS: -lphp7 -lm -ldl -lresolv -lcurl -lmcrypt -lz -lssl -lcrypto -lxml2 -lopcache
import "C"
