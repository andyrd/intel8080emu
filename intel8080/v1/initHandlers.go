package v1

import ops "github.com/andyrd/emu/intel8080"

func (v *v1) initHandlers() {
	v.handlers[ops.NOP] = v.NOP
	v.handlers[0x08] = v.NOP
	v.handlers[0x10] = v.NOP
	v.handlers[0x18] = v.NOP
	v.handlers[0x20] = v.NOP
	v.handlers[0x28] = v.NOP
	v.handlers[0x30] = v.NOP
	v.handlers[0x38] = v.NOP
	v.handlers[ops.LXI_B_D16] = v.LXI_B_D16
	v.handlers[ops.STAX_B] = v.STAX_B
	v.handlers[ops.INX_B] = v.INX_B
	v.handlers[ops.INR_B] = v.INR_B
	v.handlers[ops.DCR_B] = v.DCR_B
	v.handlers[ops.MVI_B_D8] = v.MVI_B_D8
	v.handlers[ops.RLC] = v.RLC
	v.handlers[ops.DAD_B] = v.DAD_B
	v.handlers[ops.LDAX_B] = v.LDAX_B
	v.handlers[ops.DCX_B] = v.DCX_B
	v.handlers[ops.INR_C] = v.INR_C
	v.handlers[ops.DCR_C] = v.DCR_C
	v.handlers[ops.MVI_C_D8] = v.MVI_C_D8
	v.handlers[ops.RRC] = v.RRC
	v.handlers[ops.LXI_D_D16] = v.LXI_D_D16
	v.handlers[ops.STAX_D] = v.STAX_D
	v.handlers[ops.INX_D] = v.INX_D
	v.handlers[ops.INR_D] = v.INR_D
	v.handlers[ops.DCR_D] = v.DCR_D
	v.handlers[ops.MVI_D_D8] = v.MVI_D_D8
	v.handlers[ops.RAL] = v.RAL
	v.handlers[ops.DAD_D] = v.DAD_D
	v.handlers[ops.LDAX_D] = v.LDAX_D
	v.handlers[ops.DCX_D] = v.DCX_D
	v.handlers[ops.INR_E] = v.INR_E
	v.handlers[ops.DCR_E] = v.DCR_E
	v.handlers[ops.MVI_E_D8] = v.MVI_E_D8
	v.handlers[ops.RAR] = v.RAR
	v.handlers[ops.LXI_H_D16] = v.LXI_H_D16
	v.handlers[ops.SHLD_A16] = v.SHLD_A16
	v.handlers[ops.INX_H] = v.INX_H
	v.handlers[ops.INR_H] = v.INR_H
	v.handlers[ops.DCR_H] = v.DCR_H
	v.handlers[ops.MVI_H_D8] = v.MVI_H_D8
	v.handlers[ops.DAA] = v.DAA
	v.handlers[ops.DAD_H] = v.DAD_H
	v.handlers[ops.LHLD_A16] = v.LHLD_A16
	v.handlers[ops.DCX_H] = v.DCX_H
	v.handlers[ops.INR_L] = v.INR_L
	v.handlers[ops.DCR_L] = v.DCR_L
	v.handlers[ops.MVI_L_D8] = v.MVI_L_D8
	v.handlers[ops.CMA] = v.CMA
	v.handlers[ops.LXI_SP_D16] = v.LXI_SP_D16
	v.handlers[ops.STA_A16] = v.STA_A16
	v.handlers[ops.INX_SP] = v.INX_SP
	v.handlers[ops.INR_M] = v.INR_M
	v.handlers[ops.DCR_M] = v.DCR_M
	v.handlers[ops.MVI_M_D8] = v.MVI_M_D8
	v.handlers[ops.STC] = v.STC
	v.handlers[ops.DAD_SP] = v.DAD_SP
	v.handlers[ops.LDA_A16] = v.LDA_A16
	v.handlers[ops.DCX_SP] = v.DCX_SP
	v.handlers[ops.INR_A] = v.INR_A
	v.handlers[ops.DCR_A] = v.DCR_A
	v.handlers[ops.MVI_A_D8] = v.MVI_A_D8
	v.handlers[ops.CMC] = v.CMC
	v.handlers[ops.MOV_B_B] = v.MOV_B_B
	v.handlers[ops.MOV_B_C] = v.MOV_B_C
	v.handlers[ops.MOV_B_D] = v.MOV_B_D
	v.handlers[ops.MOV_B_E] = v.MOV_B_E
	v.handlers[ops.MOV_B_H] = v.MOV_B_H
	v.handlers[ops.MOV_B_L] = v.MOV_B_L
	v.handlers[ops.MOV_B_M] = v.MOV_B_M
	v.handlers[ops.MOV_B_A] = v.MOV_B_A
	v.handlers[ops.MOV_C_B] = v.MOV_C_B
	v.handlers[ops.MOV_C_C] = v.MOV_C_C
	v.handlers[ops.MOV_C_D] = v.MOV_C_D
	v.handlers[ops.MOV_C_E] = v.MOV_C_E
	v.handlers[ops.MOV_C_H] = v.MOV_C_H
	v.handlers[ops.MOV_C_L] = v.MOV_C_L
	v.handlers[ops.MOV_C_M] = v.MOV_C_M
	v.handlers[ops.MOV_C_A] = v.MOV_C_A
	v.handlers[ops.MOV_D_B] = v.MOV_D_B
	v.handlers[ops.MOV_D_C] = v.MOV_D_C
	v.handlers[ops.MOV_D_D] = v.MOV_D_D
	v.handlers[ops.MOV_D_E] = v.MOV_D_E
	v.handlers[ops.MOV_D_H] = v.MOV_D_H
	v.handlers[ops.MOV_D_L] = v.MOV_D_L
	v.handlers[ops.MOV_D_M] = v.MOV_D_M
	v.handlers[ops.MOV_D_A] = v.MOV_D_A
	v.handlers[ops.MOV_E_B] = v.MOV_E_B
	v.handlers[ops.MOV_E_C] = v.MOV_E_C
	v.handlers[ops.MOV_E_D] = v.MOV_E_D
	v.handlers[ops.MOV_E_E] = v.MOV_E_E
	v.handlers[ops.MOV_E_H] = v.MOV_E_H
	v.handlers[ops.MOV_E_L] = v.MOV_E_L
	v.handlers[ops.MOV_E_M] = v.MOV_E_M
	v.handlers[ops.MOV_E_A] = v.MOV_E_A
	v.handlers[ops.MOV_H_B] = v.MOV_H_B
	v.handlers[ops.MOV_H_C] = v.MOV_H_C
	v.handlers[ops.MOV_H_D] = v.MOV_H_D
	v.handlers[ops.MOV_H_E] = v.MOV_H_E
	v.handlers[ops.MOV_H_H] = v.MOV_H_H
	v.handlers[ops.MOV_H_L] = v.MOV_H_L
	v.handlers[ops.MOV_H_M] = v.MOV_H_M
	v.handlers[ops.MOV_H_A] = v.MOV_H_A
	v.handlers[ops.MOV_L_B] = v.MOV_L_B
	v.handlers[ops.MOV_L_C] = v.MOV_L_C
	v.handlers[ops.MOV_L_D] = v.MOV_L_D
	v.handlers[ops.MOV_L_E] = v.MOV_L_E
	v.handlers[ops.MOV_L_H] = v.MOV_L_H
	v.handlers[ops.MOV_L_L] = v.MOV_L_L
	v.handlers[ops.MOV_L_M] = v.MOV_L_M
	v.handlers[ops.MOV_L_A] = v.MOV_L_A
	v.handlers[ops.MOV_M_B] = v.MOV_M_B
	v.handlers[ops.MOV_M_C] = v.MOV_M_C
	v.handlers[ops.MOV_M_D] = v.MOV_M_D
	v.handlers[ops.MOV_M_E] = v.MOV_M_E
	v.handlers[ops.MOV_M_H] = v.MOV_M_H
	v.handlers[ops.MOV_M_L] = v.MOV_M_L
	v.handlers[ops.HLT] = v.HLT
	v.handlers[ops.MOV_M_A] = v.MOV_M_A
	v.handlers[ops.MOV_A_B] = v.MOV_A_B
	v.handlers[ops.MOV_A_C] = v.MOV_A_C
	v.handlers[ops.MOV_A_D] = v.MOV_A_D
	v.handlers[ops.MOV_A_E] = v.MOV_A_E
	v.handlers[ops.MOV_A_H] = v.MOV_A_H
	v.handlers[ops.MOV_A_L] = v.MOV_A_L
	v.handlers[ops.MOV_A_M] = v.MOV_A_M
	v.handlers[ops.MOV_A_A] = v.MOV_A_A
	v.handlers[ops.ADD_B] = v.ADD_B
	v.handlers[ops.ADD_C] = v.ADD_C
	v.handlers[ops.ADD_D] = v.ADD_D
	v.handlers[ops.ADD_E] = v.ADD_E
	v.handlers[ops.ADD_H] = v.ADD_H
	v.handlers[ops.ADD_L] = v.ADD_L
	v.handlers[ops.ADD_M] = v.ADD_M
	v.handlers[ops.ADD_A] = v.ADD_A
	v.handlers[ops.ADC_B] = v.ADC_B
	v.handlers[ops.ADC_C] = v.ADC_C
	v.handlers[ops.ADC_D] = v.ADC_D
	v.handlers[ops.ADC_E] = v.ADC_E
	v.handlers[ops.ADC_H] = v.ADC_H
	v.handlers[ops.ADC_L] = v.ADC_L
	v.handlers[ops.ADC_M] = v.ADC_M
	v.handlers[ops.ADC_A] = v.ADC_A
	v.handlers[ops.SUB_B] = v.SUB_B
	v.handlers[ops.SUB_C] = v.SUB_C
	v.handlers[ops.SUB_D] = v.SUB_D
	v.handlers[ops.SUB_E] = v.SUB_E
	v.handlers[ops.SUB_H] = v.SUB_H
	v.handlers[ops.SUB_L] = v.SUB_L
	v.handlers[ops.SUB_M] = v.SUB_M
	v.handlers[ops.SUB_A] = v.SUB_A
	v.handlers[ops.SBB_B] = v.SBB_B
	v.handlers[ops.SBB_C] = v.SBB_C
	v.handlers[ops.SBB_D] = v.SBB_D
	v.handlers[ops.SBB_E] = v.SBB_E
	v.handlers[ops.SBB_H] = v.SBB_H
	v.handlers[ops.SBB_L] = v.SBB_L
	v.handlers[ops.SBB_M] = v.SBB_M
	v.handlers[ops.SBB_A] = v.SBB_A
}
