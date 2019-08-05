# Bitbucket Server Status

This step the build status for a commit on Bitbucket Server. It uses the Bitbucket REST API to set a build status on the current commit. By default, it will use the status "success" or "failed" based on the status of the Bitrise build. This can be overridden to provide specific statuses for other things happening in the build.

A single commit can have many build statuses. Each build status has its own `build_key`. So, you might have a `build_key` for compiling the code, and another status for successfully deploying the artifacts. 

For more information on Bitbucket build statuses, please refer to the [documentation](https://developer.atlassian.com/server/bitbucket/how-tos/updating-build-status-for-commits/).

### Inputs
* `api_base_url`: This is the URL to the REST api of the Bitbucket server where your code is hosted, e.g. https://bitbucket.example.org/rest/
* `commit_hash`: The commit to assign this status to
* `status`: This is the build status we want to assign to this commit. It can be In Progress, Successful, or Failed. If you choose "Auto", it will use the success/failure status of the current Bitrise build.
* `build_key`: This is the identifier for this status on the commit. Some examples of this might be `lint_code`, `compile`, `deploy_artifacts`. Running this step three times with those three keys would cause the commit/PR on Bitbucket to show "3 builds passed"
* `build_name`: This is the human-readable title for the build on Bitbucket that shows up under the list of builds
* `description`: The summary text that appears below the build name and link in the Bitbucket build list. 
* `target_url`: The URL to navigate to when clicking on a build in Bitbucket. Defaults to the `BITRISE_BUILD_URL` environment variable.
* `bitbucket_token`: This is the Personal Access Token that authenticates Bitrise to send REST API calls to Bitbucket. More information can be found here: https://confluence.atlassian.com/bitbucketserver/personal-access-tokens-939515499.html

### Outputs
There are no outputs from this step.

## Testing this step

Can be run directly with the [bitrise CLI](https://github.com/bitrise-io/bitrise),
just `git clone` this repository, `cd` into it's folder in your Terminal/Command Line
and call `bitrise run test`.

*Check the `bitrise.yml` file for required inputs which have to be
added to your `.bitrise.secrets.yml` file!*

Step by step:

1. Open up your Terminal / Command Line
2. `git clone` the repository
3. `cd` into the directory of the step (the one you just `git clone`d)
4. Create a `.bitrise.secrets.yml` file in the same directory of `bitrise.yml`
   (the `.bitrise.secrets.yml` is a git ignored file, you can store your secrets in it)
5. Check the `bitrise.yml` file for any secret you should set in `.bitrise.secrets.yml`
  * Best practice is to mark these options with something like `# define these in your .bitrise.secrets.yml`, in the `app:envs` section.
7. Once you have all the required secret parameters in your `.bitrise.secrets.yml` you can just run this step with the [bitrise CLI](https://github.com/bitrise-io/bitrise): `bitrise run test`

An example `.bitrise.secrets.yml` file:

```
envs:
- A_SECRET_PARAM_ONE: the value for secret one
- A_SECRET_PARAM_TWO: the value for secret two
```