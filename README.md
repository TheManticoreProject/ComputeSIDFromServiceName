![](./.github/banner.png)

<p align="center">
  A cross-platform tool to compute the value of a Windows Security Identifier (SID) from a service name.
  <br>
  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/TheManticoreProject/ComputeSIDFromServiceName">
  <a href="https://twitter.com/intent/follow?screen_name=podalirius_" title="Follow"><img src="https://img.shields.io/twitter/follow/podalirius_?label=Podalirius&style=social"></a>
  <a href="https://www.youtube.com/c/Podalirius_?sub_confirmation=1" title="Subscribe"><img alt="YouTube Channel Subscribers" src="https://img.shields.io/youtube/channel/subscribers/UCF_x5O7CSfr82AfNVTKOv_A?style=social"></a>
  <br>
</p>


## Features

- [x] Compute the value of a Windows Security Identifier (SID) from a service name

## Example

```bash
$ ./ComputeSIDFromServiceName -s "MSSQLSERVER"
S-1-5-80-1000-1000-1000-1000-1000
```

## Usage

```
$ ./ComputeSIDFromServiceName -h
Usage: ComputeSIDFromServiceName [--debug] --service-name <string>

  -d, --debug                 Debug mode. (default: false)
  -s, --service-name <string> Service name to compute SID from.
```

## Contributing

Pull requests are welcome. Feel free to open an issue if you want to add other features.

## References
 - https://pcsxcetrasupport3.wordpress.com/2013/09/08/how-do-you-get-a-service-sid-from-a-service-name/