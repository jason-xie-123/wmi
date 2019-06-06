﻿// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using Microsoft.WmiCodeGen.CSharp;

namespace Microsoft.WmiCodeGen
{
    public class GoFormat : IFormat
    {
        #region Constructor
        public GoFormat(string[] wmiNamespaces, string outDir, bool recurse) : base(wmiNamespaces, outDir, recurse)
        {
        }
        #endregion

        protected override WmiNamespace GetWmiNamespace(string wmins)
        {
            return new CSWmiNamespace(wmins);
        }
    }
}