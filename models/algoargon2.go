package models


// AlgoArgon2 Model
type AlgoArgon2 struct {
    // Algo type.
    Type string `json:"xtype"`
    // Memory used to compute hash.
    MemoryCost int `json:"memoryCost"`
    // Amount of time consumed to compute hash
    TimeCost int `json:"timeCost"`
    // Number of threads used to compute hash.
    Threads int `json:"threads"`

}
