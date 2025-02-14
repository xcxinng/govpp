//  Copyright (c) 2020 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package binapigen

func init() {
	//RegisterPlugin("convert", GenerateConvert)
}

// library dependencies
const (
	fmtPkg     = GoImportPath("fmt")
	netPkg     = GoImportPath("net")
	timePkg    = GoImportPath("time")
	stringsPkg = GoImportPath("strings")
)

func genIPConversion(g *GenFile, structName string, ipv int) {
	// ParseIPXAddress method
	g.P("func Parse", structName, "(s string) (", structName, ", error) {")
	if ipv == 4 {
		g.P("	ip := ", netPkg.Ident("ParseIP"), "(s).To4()")
	} else {
		g.P("	ip := ", netPkg.Ident("ParseIP"), "(s).To16()")
	}
	g.P("	if ip == nil {")
	g.P("		return ", structName, "{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP address: %s\", s)")
	g.P("	}")
	g.P("	var ipaddr ", structName)
	if ipv == 4 {
		g.P("	copy(ipaddr[:], ip.To4())")
	} else {
		g.P("	copy(ipaddr[:], ip.To16())")
	}
	g.P("	return ipaddr, nil")
	g.P("}")
	g.P()

	// ToIP method
	g.P("func (x ", structName, ") ToIP() ", netPkg.Ident("IP"), " {")
	if ipv == 4 {
		g.P("	return ", netPkg.Ident("IP"), "(x[:]).To4()")
	} else {
		g.P("	return ", netPkg.Ident("IP"), "(x[:]).To16()")
	}
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	return x.ToIP().String()")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.String()), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	ipaddr, err := Parse", structName, "(string(text))")
	g.P("	if err !=nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = ipaddr")
	g.P("	return nil")
	g.P("}")
	g.P()
}

func genAddressConversion(g *GenFile, structName string) {
	// ParseAddress method
	g.P("func Parse", structName, "(s string) (", structName, ", error) {")
	g.P("	ip := ", netPkg.Ident("ParseIP"), "(s)")
	g.P("	if ip == nil {")
	g.P("		return ", structName, "{}, ", fmtPkg.Ident("Errorf"), "(\"invalid address: %s\", s)")
	g.P("	}")
	g.P("	return ", structName, "FromIP(ip), nil")
	g.P("}")
	g.P()

	// AddressFromIP method
	g.P("func ", structName, "FromIP(ip ", netPkg.Ident("IP"), ") ", structName, " {")
	g.P("	var addr ", structName)
	g.P("	if ip.To4() == nil {")
	g.P("		addr.Af = ADDRESS_IP6")
	g.P("		var ip6 IP6Address")
	g.P("		copy(ip6[:], ip.To16())")
	g.P("		addr.Un.SetIP6(ip6)")
	g.P("	} else {")
	g.P("		addr.Af = ADDRESS_IP4")
	g.P("		var ip4 IP4Address")
	g.P("		copy(ip4[:], ip.To4())")
	g.P("		addr.Un.SetIP4(ip4)")
	g.P("	}")
	g.P("	return addr")
	g.P("}")
	g.P()

	// ToIP method
	g.P("func (x ", structName, ") ToIP() ", netPkg.Ident("IP"), " {")
	g.P("	if x.Af == ADDRESS_IP6 {")
	g.P("		ip6 := x.Un.GetIP6()")
	g.P("		return ", netPkg.Ident("IP"), "(ip6[:]).To16()")
	g.P("	} else {")
	g.P("		ip4 := x.Un.GetIP4()")
	g.P("		return ", netPkg.Ident("IP"), "(ip4[:]).To4()")
	g.P("	}")
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	return x.ToIP().String()")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.String()), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	addr, err := Parse", structName, "(string(text))")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = addr")
	g.P("	return nil")
	g.P("}")
	g.P()
}

func genIPPrefixConversion(g *GenFile, structName string, ipv int) {
	// ParsePrefix method
	g.P("func Parse", structName, "(s string) (prefix ", structName, ", err error) {")
	g.P("	hasPrefix := ", stringsPkg.Ident("Contains"), "(s, \"/\")")
	g.P("	if hasPrefix {")
	g.P("		ip, network, err := ", netPkg.Ident("ParseCIDR"), "(s)")
	g.P("		if err != nil {")
	g.P("			return ", structName, "{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP %s: %s\", s, err)")
	g.P("		}")
	g.P("		maskSize, _ := network.Mask.Size()")
	g.P("		prefix.Len = byte(maskSize)")
	if ipv == 4 {
		g.P("		prefix.Address, err = ParseIP4Address(ip.String())")
	} else {
		g.P("		prefix.Address, err = ParseIP6Address(ip.String())")
	}
	g.P("		if err != nil {")
	g.P("			return ", structName, "{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP %s: %s\", s, err)")
	g.P("		}")
	g.P("	} else {")
	g.P("		ip :=  ", netPkg.Ident("ParseIP"), "(s)")
	g.P("		defaultMaskSize, _ := ", netPkg.Ident("CIDRMask"), "(32, 32).Size()")
	g.P("		if ip.To4() == nil {")
	g.P("			defaultMaskSize, _ =", netPkg.Ident("CIDRMask"), "(128, 128).Size()")
	g.P("		}")
	g.P("		prefix.Len = byte(defaultMaskSize)")
	if ipv == 4 {
		g.P("		prefix.Address, err = ParseIP4Address(ip.String())")
	} else {
		g.P("		prefix.Address, err = ParseIP6Address(ip.String())")
	}
	g.P("		if err != nil {")
	g.P("			return ", structName, "{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP %s: %s\", s, err)")
	g.P("		}")
	g.P("	}")
	g.P("	return prefix, nil")
	g.P("}")
	g.P()

	// ToIPNet method
	g.P("func (x ", structName, ") ToIPNet() *", netPkg.Ident("IPNet"), " {")
	if ipv == 4 {
		g.P("	mask := ", netPkg.Ident("CIDRMask"), "(int(x.Len), 32)")
	} else {
		g.P("	mask := ", netPkg.Ident("CIDRMask"), "(int(x.Len), 128)")
	}
	g.P("	ipnet := &", netPkg.Ident("IPNet"), "{IP: x.Address.ToIP(), Mask: mask}")
	g.P("	return ipnet")
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	ip := x.Address.String()")
	g.P("	return ip + \"/\" + ", strconvPkg.Ident("Itoa"), "(int(x.Len))")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.String()), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	prefix, err := Parse", structName, "(string(text))")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = prefix")
	g.P("	return nil")
	g.P("}")
	g.P()
}

func genPrefixConversion(g *GenFile, structName string) {
	// ParsePrefix method
	g.P("func Parse", structName, "(ip string) (prefix ", structName, ", err error) {")
	g.P("	hasPrefix := ", stringsPkg.Ident("Contains"), "(ip, \"/\")")
	g.P("	if hasPrefix {")
	g.P("		netIP, network, err := ", netPkg.Ident("ParseCIDR"), "(ip)")
	g.P("		if err != nil {")
	g.P("			return Prefix{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP %s: %s\", ip, err)")
	g.P("		}")
	g.P("		maskSize, _ := network.Mask.Size()")
	g.P("		prefix.Len = byte(maskSize)")
	g.P("		prefix.Address, err = ParseAddress(netIP.String())")
	g.P("		if err != nil {")
	g.P("			return Prefix{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP %s: %s\", ip, err)")
	g.P("		}")
	g.P("	} else {")
	g.P("		netIP :=  ", netPkg.Ident("ParseIP"), "(ip)")
	g.P("		defaultMaskSize, _ := ", netPkg.Ident("CIDRMask"), "(32, 32).Size()")
	g.P("		if netIP.To4() == nil {")
	g.P("			defaultMaskSize, _ =", netPkg.Ident("CIDRMask"), "(128, 128).Size()")
	g.P("		}")
	g.P("		prefix.Len = byte(defaultMaskSize)")
	g.P("		prefix.Address, err = ParseAddress(netIP.String())")
	g.P("		if err != nil {")
	g.P("			return Prefix{}, ", fmtPkg.Ident("Errorf"), "(\"invalid IP %s: %s\", ip, err)")
	g.P("		}")
	g.P("	}")
	g.P("	return prefix, nil")
	g.P("}")
	g.P()

	// ToIPNet method
	g.P("func (x ", structName, ") ToIPNet() *", netPkg.Ident("IPNet"), " {")
	g.P("	var mask ", netPkg.Ident("IPMask"))
	g.P("	if x.Address.Af == ADDRESS_IP4 {")
	g.P("	mask = ", netPkg.Ident("CIDRMask"), "(int(x.Len), 32)")
	g.P("	} else {")
	g.P("	mask = ", netPkg.Ident("CIDRMask"), "(int(x.Len), 128)")
	g.P("	}")
	g.P("	ipnet := &", netPkg.Ident("IPNet"), "{IP: x.Address.ToIP(), Mask: mask}")
	g.P("	return ipnet")
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	ip := x.Address.String()")
	g.P("	return ip + \"/\" + ", strconvPkg.Ident("Itoa"), "(int(x.Len))")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.String()), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	prefix, err := Parse", structName, "(string(text))")
	g.P("	if err !=nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = prefix")
	g.P("	return nil")
	g.P("}")
	g.P()
}

func genAddressWithPrefixConversion(g *GenFile, structName string) {
	// ParseAddressWithPrefix method
	g.P("func Parse", structName, "(s string) (", structName, ", error) {")
	g.P("	prefix, err := ParsePrefix(s)")
	g.P("	if err != nil {")
	g.P("		return ", structName, "{}, err")
	g.P("	}")
	g.P("	return ", structName, "(prefix), nil")
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	return Prefix(x).String()")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.String()), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	prefix, err := Parse", structName, "(string(text))")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = prefix")
	g.P("	return nil")
	g.P("}")
	g.P()
}

func genMacAddressConversion(g *GenFile, structName string) {
	// ParseMAC method
	g.P("func Parse", structName, "(s string) (", structName, ", error) {")
	g.P("	var macaddr ", structName)
	g.P("	mac, err := ", netPkg.Ident("ParseMAC"), "(s)")
	g.P("	if err != nil {")
	g.P("		return macaddr, err")
	g.P("	}")
	g.P("	copy(macaddr[:], mac[:])")
	g.P("	return macaddr, nil")
	g.P("}")
	g.P()

	// ToMAC method
	g.P("func (x ", structName, ") ToMAC() ", netPkg.Ident("HardwareAddr"), " {")
	g.P("	return ", netPkg.Ident("HardwareAddr"), "(x[:])")
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	return x.ToMAC().String()")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.String()), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	mac, err := Parse", structName, "(string(text))")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = mac")
	g.P("	return nil")
	g.P("}")
	g.P()
}

func genTimestampConversion(g *GenFile, structName string) {
	// NewTimestamp method
	g.P("func New", structName, "(t ", timePkg.Ident("Time"), ") ", structName, " {")
	g.P("	sec := int64(t.Unix())")
	g.P("	nsec := int32(t.Nanosecond())")
	g.P("	ns := float64(sec) + float64(nsec / 1e9)")
	g.P("	return ", structName, "(ns)")
	g.P("}")
	g.P()

	// ToTime method
	g.P("func (x ", structName, ") ToTime() ", timePkg.Ident("Time"), " {")
	g.P("	ns := int64(x * 1e9)")
	g.P("	sec := ns / 1e9")
	g.P("	nsec := ns % 1e9")
	g.P("	return ", timePkg.Ident("Unix"), "(sec, nsec)")
	g.P("}")
	g.P()

	// String method
	g.P("func (x ", structName, ") String() string {")
	g.P("	return x.ToTime().String()")
	g.P("}")
	g.P()

	// MarshalText method
	g.P("func (x *", structName, ") MarshalText() ([]byte, error) {")
	g.P("	return []byte(x.ToTime().Format(", timePkg.Ident("RFC3339Nano"), ")), nil")
	g.P("}")
	g.P()

	// UnmarshalText method
	g.P("func (x *", structName, ") UnmarshalText(text []byte) error {")
	g.P("	t, err := ", timePkg.Ident("Parse"), "(", timePkg.Ident("RFC3339Nano"), ", string(text))")
	g.P("	if err != nil {")
	g.P("		return err")
	g.P("	}")
	g.P("	*x = New", structName, "(t)")
	g.P("	return nil")
	g.P("}")
	g.P()
}
