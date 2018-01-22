package supply

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack"
)

type Stager interface {
	LinkDirectoryInDepDir(string, string) error
	DepDir() string
}

type Manifest interface {
	InstallOnlyVersion(depName string, installDir string) error
}

type Supplier struct {
	Stager   Stager
	Manifest Manifest
	Log      *libbuildpack.Logger
}

func New(stager Stager, manifest Manifest, logger *libbuildpack.Logger) *Supplier {
	return &Supplier{
		Stager:   stager,
		Manifest: manifest,
		Log:      logger,
	}
}

func (s *Supplier) Run() error {
	s.Log.BeginStep("Supplying JDK")

	if err := s.Manifest.InstallOnlyVersion("openjdk", filepath.Join(s.Stager.DepDir(), "openjdk")); err != nil {
		s.Log.Error("Could not install JDK: %s", err)
		return err
	}

	for _, dir := range []string{"bin", "lib"} {
		if err := s.Stager.LinkDirectoryInDepDir(filepath.Join(s.Stager.DepDir(), "openjdk", dir), dir); err != nil {
			s.Log.Error("Could not link JDK/%s: %s", dir, err)
			return err
		}
	}

	return nil
}
