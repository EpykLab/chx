package output

// alienvault
func (s *AlienVaultDomain) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(AVDOMAIN, d)
	return nil
}
