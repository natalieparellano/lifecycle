package main

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/buildpacks/lifecycle/cmd"
	"github.com/buildpacks/lifecycle/env"
	"github.com/buildpacks/lifecycle/launch"
)

func main() {
	cmd.Exit(runLaunch())
}

func runLaunch() error {
	platformAPI := cmd.EnvOrDefault(cmd.EnvPlatformAPI, cmd.DefaultPlatformAPI)
	if err := cmd.VerifyPlatformAPI(platformAPI); err != nil {
		cmd.Exit(err)
	}

	var md launch.Metadata
	if _, err := toml.DecodeFile(launch.GetMetadataFilePath(cmd.EnvOrDefault(cmd.EnvLayersDir, cmd.DefaultLayersDir)), &md); err != nil {
		return cmd.FailErr(err, "read metadata")
	}
	if err := verifyBuildpackApis(md.Buildpacks); err != nil {
		return err
	}

	launcher := &launch.Launcher{
		DefaultProcessType: cmd.EnvOrDefault(cmd.EnvProcessType, cmd.DefaultProcessType),
		LayersDir:          cmd.EnvOrDefault(cmd.EnvLayersDir, cmd.DefaultLayersDir),
		AppDir:             cmd.EnvOrDefault(cmd.EnvAppDir, cmd.DefaultAppDir),
		Processes:          md.Processes,
		Buildpacks:         md.Buildpacks,
		Env:                env.NewLaunchEnv(os.Environ()),
		Exec:               launch.OSExecFunc,
		Setenv:             os.Setenv,
	}

	if err := launcher.Launch(os.Args[0], os.Args[1:]); err != nil {
		return cmd.FailErrCode(err, cmd.CodeFailedLaunch, "launch")
	}
	return nil
}

func verifyBuildpackApis(bps []launch.Buildpack) error {
	for _, bp := range bps {
		if bp.API == "" {
			// If the same lifecycle is used for build and launcher we should never end up here
			// but if for some reason we do, default to 0.2
			bp.API = "0.2"
		}
		if err := cmd.VerifyBuildpackAPI(bp.ID, bp.API); err != nil {
			return err
		}
	}
	return nil
}
