package domain

type Patch struct {
  ID  string `json:"id"`
  Title string `json:"title"`
  AuthorName string `json:"authorName"`
  AuthorEmail string `json:"authorEmail"`
  PatchHref string `json:"patchHref"`
  PatchTag string `json:"patchTag"`
}
