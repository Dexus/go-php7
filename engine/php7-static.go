// Copyright 2016 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package engine

/* 
There is no official PHP 7.1 in the Ubuntu 16.04 repos.

If you want PHP 7.1, there is a version available in ppa:ondrej/php

You can install it like this:

> sudo add-apt-repository ppa:ondrej/php
> sudo apt-get update
> sudo apt-get remove php7.0 #(optional) 
> sudo apt-get install php7.1 #(from comments)

Remember that this is not an official upgrade path. The PPA is well known, and is relatively safe to use.
*/

// #cgo CFLAGS: -I/usr/include/php/20160303 -Iinclude/php7 -Isrc/php7 -Iinclude
// #cgo CFLAGS: -I/usr/include/php/20160303 -I/usr/include/php/20160303/main -I/usr/include/php/20160303/TSRM -I/usr/include/php/20160303/Zend
// #cgo CFLAGS: -I/usr/include/php/20160303/ext -I/usr/include/php/20160303/ext/date/lib
// #cgo LDFLAGS: -L/usr/lib/php/20160303  -L/usr/lib/x86_64-linux-gnu
// #cgo LDFLAGS: -lphp7.1 -lm -ldl -lcrypt -lresolv -lcrypt -lz -lpcre -lrt -lm -ldl -lnsl -lcurl -lxml2 -lssl -lcrypto -lcrypt -lcrypt -lmcrypt -lopcache
import "C"
