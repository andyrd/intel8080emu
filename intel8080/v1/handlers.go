package v1

func (v *v1) NOP() {
	v.cycles -= 4
}

func (v *v1) LXI_B_D16() {
	v.state.C = v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.B = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 10
}

func (v *v1) STAX_B() {
	v.state.Memory[uint16(v.state.B)<<8|uint16(v.state.C)] = v.state.A
	v.cycles -= 7
}

func (v *v1) INX_B() {
	result := (uint16(v.state.B)<<8 | uint16(v.state.C)) + 1
	v.state.B = byte(result >> 8)
	v.state.C = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_B() {
	result := v.state.B + 1
	v.setFlag(halfCarryFlagBit, halfCarryAdd(v.state.B, 1))
	v.setFlag(parityFlagBit, parity(result))
	v.setFlag(zeroFlagBit, zero(result))
	v.setFlag(signFlagBit, sign(result))
	v.state.B = result
	v.cycles -= 5
}

func (v *v1) DCR_B() {
	result := v.state.B - 1
	v.setFlag(halfCarryFlagBit, halfCarrySub(v.state.B, 1))
	v.setFlag(parityFlagBit, parity(result))
	v.setFlag(zeroFlagBit, zero(result))
	v.setFlag(signFlagBit, sign(result))
	v.state.B = result
	v.cycles -= 5
}