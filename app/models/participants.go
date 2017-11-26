package models

type Participant struct {
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	Blacklist []*Receiver `json:"blacklist"`
	*Receiver `json:"-"`
}

func (r *Participant) GetBlacklistNames() []string {
	var blacklistNames []string
	for _, blacklister := range r.Blacklist {
		blacklistNames = append(blacklistNames, blacklister.Name)
	}
	return blacklistNames
}

type Receiver struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
