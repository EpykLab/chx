package pretty

func (s VirusTotalHash) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(VT, d)
	return nil
}

func (s AlienVaultHash) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(AVHASH, d)
	return nil
}
