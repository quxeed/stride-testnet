package types

// // DefaultGenesis returns the default Capability genesis state
func NewDepositRecord(amount int64, denom string, hostZoneId string, status DepositRecord_Status, epochNumber uint64) *DepositRecord {
	return &DepositRecord{
		Id:         0,
		Amount:     amount,
		Denom:      denom,
		HostZoneId: hostZoneId,
		Status:    status,
		EpochNumber: epochNumber,
	}
}
