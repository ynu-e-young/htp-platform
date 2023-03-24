<div id="top"></div>

<!-- PROJECT SHIELDS -->
<p align="center">
<a href="https://github.com/ynu-e-young/htp-platform/graphs/contributors"><img src="https://img.shields.io/github/contributors/ynu-e-young/htp-platform.svg?style=for-the-badge" alt="Contributors"></a>
<a href="https://github.com/ynu-e-young/htp-platform/network/members"><img src="https://img.shields.io/github/forks/ynu-e-young/htp-platform.svg?style=for-the-badge" alt="Forks"></a>
<a href="https://github.com/ynu-e-young/htp-platform/stargazers"><img src="https://img.shields.io/github/stars/ynu-e-young/htp-platform.svg?style=for-the-badge" alt="Stargazers"></a>
<a href="https://github.com/ynu-e-young/htp-platform/issues"><img src="https://img.shields.io/github/issues/ynu-e-young/htp-platform.svg?style=for-the-badge" alt="Issues"></a>
<a href="https://github.com/ynu-e-young/htp-platform/actions/workflows/docker-publish.yml"><img src="https://img.shields.io/github/actions/workflow/status/ynu-e-young/htp-platform/docker-publish.yml?branch=main&style=for-the-badge" alt="Deploy"></a>
</p>

<!-- PROJECT LOGO -->
<br/>
<div align="center">
<!--   <a href="https://github.com/ynu-e-young/htp-platform">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">htp-platform</h3>

  <p align="center">
    高附加值作物的智能表型测定系统——基于新型绳驱并联机器人
    <br/>
    <a href="https://hominsu.github.io/htp-platform/"><strong>Explore the docs »</strong></a>
    <br/>
    <br/>
    <a href="https://github.com/ynu-e-young/htp-platform">View Demo</a>
    ·
    <a href="https://github.com/ynu-e-young/htp-platform/issues">Report Bug</a>
    ·
    <a href="https://github.com/ynu-e-young/htp-platform/issues">Request Feature</a>
  </p>
</div>

## Description

高附加值作物的智能表型测定系统——基于新型绳驱并联机器人

## Details

```mermaid
flowchart LR
	interface("interface service") <-.-> machine("machine service")
	interface <-.-> user("user service")
	interface <-.images.-> capture("capture service")
	machine("machine service") <-.images.-> capture
	machine("machine service") <-.-> robot("robot")
	
	subgraph Data Base
	machinedb[("machine db")]
	userdb[("user db")]
	end
	
	subgraph OSS
	ali-oss(("ali-oss"))
	end
	
	user <-.user info.-> userdb
	machine <-.machine info and log.-> machinedb
	machine -.images.-> ali-oss
	
```

