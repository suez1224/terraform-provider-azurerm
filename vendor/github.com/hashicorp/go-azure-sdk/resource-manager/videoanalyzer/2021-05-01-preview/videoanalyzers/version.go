package videoanalyzers

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2021-05-01-preview"

func userAgent() string {
	return fmt.Sprintf("hashicorp/go-azure-sdk/videoanalyzers/%s", defaultApiVersion)
}