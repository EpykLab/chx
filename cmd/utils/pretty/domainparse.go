package pretty

// alienvault
func (s *AlienVaultDomain) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(AV, d)
	return nil
}
