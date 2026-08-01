# Releasing

A published version is immutable. The registry records the checksums a version had when
it was first indexed and serves those permanently, so republishing a version with
different bytes breaks every install of it. Never re-tag and never rebuild a published
version — release the next patch version instead.

## Steps

1. Land the change on `full-device-coverage`.
2. Add the `CHANGELOG.md` entry for the new version. The release will not build without
   one.
3. Tag:

   ```bash
   git fetch origin
   git tag -a v2.0.4 origin/full-device-coverage -m "v2.0.4"
   git push origin v2.0.4
   ```

4. The release workflow validates the tag, re-runs CI against it, builds every platform,
   signs the checksums and leaves a **draft** release.
5. Review the draft at
   <https://github.com/ripclap/terraform-provider-routeros/releases>. To change anything,
   delete the draft and re-run the workflow.
6. Publish the draft. This step is irreversible.
7. `Verify published release` then confirms the registry serves the same checksums as the
   release.

## Preconditions checked before a build

- the tag is `vMAJOR.MINOR.PATCH` and is an ancestor of `full-device-coverage`
- the version is not already in the registry or published on GitHub
- `CHANGELOG.md` has an entry and the registry manifest parses

## Repository configuration

Secrets, under Settings → Secrets and variables → Actions:

| name | value |
|---|---|
| `GPG_PRIVATE_KEY` | `gpg --armor --export-secret-keys <fingerprint>` |
| `GPG_PASSPHRASE` | its passphrase |

Variable:

| name | value |
|---|---|
| `GPG_FINGERPRINT` | the fingerprint releases must be signed with |

The public key must be registered with the OpenTofu Registry for this namespace.

## Local builds

Release only from CI. Locally, build snapshots:

```bash
goreleaser release --snapshot --clean --skip=publish,sign
```

CI also retains snapshot archives from each push under the run's Artifacts.
