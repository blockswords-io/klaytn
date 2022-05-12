package governance

// Compile-time type implementation check
var _ HeaderEngine = (*Governance)(nil)
var _ ReaderEngine = (*ContractEngine)(nil)
var _ Engine = (*MixedEngine)(nil)
