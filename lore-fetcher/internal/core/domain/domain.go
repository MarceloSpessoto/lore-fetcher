package domain

type Patch struct {
  ID  string `json:"id"`
  Title string `json:"title"`
  AuthorName string `json:"authorName"`
  AuthorEmail string `json:"authorEmail"`
  PatchHref string `json:"patchHref"`
  PatchTag string `json:"patchTag"`
}

type Job struct {
	ID string `json:"id"`
	PatchID string `json:"patchId"`
	KernelArtifact []byte `json:"kernelArtifact"`
	Status string `json:"status"`
	Log string `json:"log"`
}

type dut struct {
	ID string `json:"id"`
}

type CitronInstance struct {
	ID string `json:"id"`
	GatewayIP string `json:"gatewayIp"`
	Duts []dut `json:"duts"`
}
