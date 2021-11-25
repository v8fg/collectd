# **[Cgroups](https://zh.wikipedia.org/wiki/Cgroups)**

>If app run in docker or k8s, pid shall equal **1** in normally.

## k8s

```bash
11:blkio:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
10:freezer:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
9:pids:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
8:net_prio,net_cls:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
7:perf_event:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
6:cpuacct,cpu:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
5:hugetlb:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
4:memory:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
3:devices:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
2:cpuset:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
1:name=systemd:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod5bd00f86_afbb_48be_8c28_870231f0c8e8.slice/docker-66ab041f4941037ff8067b739df85fa85a31ef383f9efccb6890b98847cfceda.scope
```

## docker

>docker in CentOS

```bash
11:blkio:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
10:perf_event:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
9:freezer:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
8:hugetlb:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
7:memory:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
6:cpuset:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
5:cpuacct,cpu:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
4:pids:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
3:devices:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
2:net_prio,net_cls:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
1:name=systemd:/docker/a8b7e600c12c19eda5b65625cfc91636d0f96248dea374d59387e813918333be
```

## Linux

```bash
11:blkio:/
10:perf_event:/
9:freezer:/
8:hugetlb:/
7:memory:/
6:cpuset:/
5:cpuacct,cpu:/
4:pids:/
3:devices:/
2:net_prio,net_cls:/
1:name=systemd:/
```
