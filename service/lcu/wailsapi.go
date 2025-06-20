package lcu


// WailsAPI exposes methods for the frontend via Wails service
// to interact with the local League client.
type WailsAPI struct{}

func (WailsAPI) ID() string { return "LcuApi" }

type AuthInfo struct {
    Port  int    `json:"port"`
    Token string `json:"token"`
}

type SummonerInfo struct {
    DisplayName  string `json:"displayName"`
    ProfileIconId int   `json:"profileIconId"`
    Region       string `json:"region"`
}

// GetAuthInfo retrieves the LCU authentication information.
func (WailsAPI) GetAuthInfo() (AuthInfo, error) {
    port, token, err := GetLolClientApiInfo()
    if err != nil {
        return AuthInfo{}, err
    }
    return AuthInfo{Port: port, Token: token}, nil
}

// GetCurrentSummoner returns basic info about the current summoner.
func (WailsAPI) GetCurrentSummoner() (SummonerInfo, error) {
    summoner, err := GetCurrSummoner()
    if err != nil {
        return SummonerInfo{}, err
    }
    info := SummonerInfo{
        DisplayName:  summoner.DisplayName,
        ProfileIconId: summoner.ProfileIconId,
        Region:       summoner.TagLine,
    }
    return info, nil
}

// StartClient attempts to start the League client if required.
// This is a stub implementation.
func (WailsAPI) StartClient() error {
    return nil
}

