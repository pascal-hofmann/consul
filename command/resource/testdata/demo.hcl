# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1

ID {
  Type = gvk("demo.v2.Artist")
  Name = "korn"
  Tenancy {
    Namespace = "default"
    Partition = "default"
  }
}

Data {
  Name = "Korn"
  Genre = "GENRE_METAL"
}

Metadata = {
  "foo" = "bar"
}
