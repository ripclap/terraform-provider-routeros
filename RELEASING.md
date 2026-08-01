# Releasing

## A published version is immutable

The registry records the checksums a version had when it was first indexed and serves
them permanently. Upload different bytes under a version that has already been published
and every install of it fails with

```
registry response indicates a package of size N, but received a package of size M
```

There is no way to correct this short of asking the registry maintainers to remove the
version. Version 2.0.2 of this provider was lost that way.

So: never re-tag, and never rebuild a published version. Release the next patch version
and mark the broken one withdrawn in `README.md`. Drafts exist precisely so a build can
be discarded and repeated before any of this becomes permanent.

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
- the signing key matches the `GPG_FINGERPRINT` variable

## What a release ships

| artefact | purpose |
|---|---|
| 13 platform archives | the provider itself |
| `*_SHA256SUMS` + `.sig` | what the registry verifies, signed with the namespace key |
| `*_manifest.json` | the registry protocol manifest |
| `*.sbom.json` | CycloneDX inventory, one per archive |
| provenance attestation | proves which workflow and commit built the archives |

The draft is checked before it can be published: every checksum is re-verified against
the files actually uploaded, each platform archive is confirmed present, and the archive
and SBOM counts must agree.

Consumers verify provenance with:

```bash
gh attestation verify <archive> --repo ripclap/terraform-provider-routeros
```

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
make snapshot
```

That skips publishing, signing and SBOM generation, so neither the signing key nor syft
is needed. CI also retains snapshot archives from each push under the run's Artifacts.
