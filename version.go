package newt

import (
	"strings"
)

const (
    // Version number of release
    Version = "0.0.1"

    // ReleaseDate, the date version.go was generated
    ReleaseDate = "2023-06-05"

    // ReleaseHash, the Git hash when version.go was generated
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
    ReleaseHash = "cdae4cb"
=======
    ReleaseHash = "c3db5c0"
>>>>>>> 13a8c0bc414fc612da91eba1fbacfc8b8973ed13
=======
    ReleaseHash = "8320ad8"
>>>>>>> 9d4adca9f257e614563efd08a6c971af97f19b49
=======
    ReleaseHash = "e374317"
>>>>>>> 763ec1a429ed0a9326f33fe7e3edff6a64c4ccf1
=======
    ReleaseHash = "763ec1a"
>>>>>>> c6dd4b0977b25c5781ac0fd5eb02feac3139f013
=======
    ReleaseHash = "c6dd4b0"
>>>>>>> f1a26be7d0615f45f4c86266afd9dc8e5d610c08
=======
    ReleaseHash = "7102725"
>>>>>>> 79065196b660cd103695ef9e101ea6fcc73231fc

    LicenseText = `
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

1. Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A
PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`
)

// FmtHelp lets you process a text block with simple curly brace markup.
func FmtHelp(src string, appName string, version string, releaseDate string, releaseHash string) string {
	m := map[string]string {
		"{app_name}": appName,
		"{version}": version,
		"{release_date}": releaseDate,
		"{release_hash}": releaseHash,
	}
	for k, v := range m {
		if strings.Contains(src, k) {
			src = strings.ReplaceAll(src, k, v)
		}
	}
	return src
}
