package models


// AlgoScrypt Model
type AlgoScrypt struct {
    // Algo type.
    Type string `json:"xtype"`
    // CPU complexity of computed hash.
    CostCpu int `json:"costCpu"`
    // Memory complexity of computed hash.
    CostMemory int `json:"costMemory"`
    // Parallelization of computed hash.
    CostParallel int `json:"costParallel"`
    // Length used to compute hash.
    Length int `json:"length"`

}
