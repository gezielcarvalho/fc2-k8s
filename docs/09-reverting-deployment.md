# How to Revert a Deployment

In Kubernetes, you can revert a Deployment to a previous state using the `kubectl rollout undo` command. This is useful if you need to roll back to a previous version of your application due to issues with the current version.

## Step 1: Check the Current Deployment
First, check the current state of your Deployment to see its history and the current version:

```bash
kubectl get deployment <deployment-name> -n <namespace>
```
Replace `<deployment-name>` with the name of your Deployment and `<namespace>` with the namespace where it is deployed.
## Step 2: View the Rollout History
You can view the rollout history of your Deployment to see the revisions available for rollback:

```bash
kubectl rollout history deployment/<deployment-name> -n <namespace>
```
This command will show you a list of revisions along with their change-cause annotations, if any.
## Step 3: Rollback to a Previous Revision
To revert to a previous revision, use the following command:

```bash
kubectl rollout undo deployment/<deployment-name> --to-revision=<revision-number> -n <namespace>
```
Replace `<revision-number>` with the revision number you want to roll back to. You can also omit the `--to-revision` flag to roll back to the previous revision automatically.
## Step 4: Verify the Rollback
After performing the rollback, check the status of your Deployment to ensure it has been reverted successfully:

```bash
kubectl get deployment <deployment-name> -n <namespace>
```
You can also check the rollout status to see if the rollback was successful:

```bash
kubectl rollout status deployment/<deployment-name> -n <namespace>
```
## Step 5: Check the Application
Finally, verify that your application is running as expected after the rollback. You can check the logs of the pods or access the application to ensure everything is functioning correctly.

```bash
kubectl logs <pod-name> -n <namespace>
```
Replace `<pod-name>` with the name of one of the pods created by your Deployment.
## Conclusion
Rolling back a Deployment in Kubernetes is a straightforward process using the `kubectl rollout undo` command. By following the steps outlined above, you can easily revert to a previous version of your application if needed.
