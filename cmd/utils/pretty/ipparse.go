package pretty

func (s AlientVaultIP) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(AVIP, d)
	return nil
}

func (s CrowdSecIP) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(CS, d)
	return nil
}

func (s AipdbIP) printer(d any) error {
	t, err := New()
	if err != nil {
		return err
	}

	t.Render(AIP, d)
	return nil
}
