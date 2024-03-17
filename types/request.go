package types

type Resource struct {
    Embedded   Embedded `json:"_embedded"`
    TotalItems int      `json:"total_items"`
}

type Embedded struct {
    Entries []Entry `json:"_entries"`
}

type Entry struct {
    Links Links `json:"_links"`
}

type Links map[string][]Link

type Link struct {
    Href string `json:"href"`
}
