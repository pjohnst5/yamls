switch {
	case strings.Contains(windowsSku, "2004"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:windowsservercore-2004",
			ServerCore: "mcr.microsoft.com/windows/servercore:2004"}, nil
	case strings.Contains(windowsSku, "1909"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:windowsservercore-1909",
			ServerCore: "mcr.microsoft.com/windows/servercore:1909"}, nil
	case strings.Contains(windowsSku, "1903"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:windowsservercore-1903",
			ServerCore: "mcr.microsoft.com/windows/servercore:1903"}, nil
	case strings.Contains(windowsSku, "1809"), strings.Contains(windowsSku, "2019"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:20191112-windowsservercore-ltsc2019",
			ServerCore: "mcr.microsoft.com/windows/servercore:ltsc2019"}, nil
	case strings.Contains(windowsSku, "1803"):
		return nil, errors.New("Windows Server version 1803 is out of support")
	case strings.Contains(windowsSku, "1709"):
		return nil, errors.New("Windows Server version 1709 is out of support")
  }
  
  mcr.microsoft.com/dotnet/framework/samples:aspnetapp