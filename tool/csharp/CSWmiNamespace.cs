﻿// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen.CSharp
{
    public class CSWmiNamespace : WmiNamespace
    {
        public CSWmiNamespace(string name) : base(name)
        {
        }

        protected override WmiReference GetReference(string reference)
        {
            return new CSWmiReference(reference);
        }

        /// <summary>
        /// Generates Sources files
        /// </summary>
        /// <param name="key"></param>
        /// <param name="valueList"></param>
        /// <param name="outDirectory"></param>
        protected override void GenerateSourceFile(string sourceFile)
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendFormat(CultureInfo.InvariantCulture,
@"####################################################################
#
#       File:           sources
#
#       Contains:       This file contains necessary information needed
#                       for {0} to build
#
#       Written by:     Auto Generated (MadhanM@Microsoft.com)
#
#       Copyright:      {1} Microsoft Corporation
#
####################################################################
MANAGED_CODE=1
URT_VER=4.5
#CLSCompliant   = true
METADATA_DLL    = metadata_dll
ComVisible      = false
TARGETTYPE      = DYNLINK
TARGET_DESTINATION  = vmtest\common\managed\wmi

ASSEMBLY_IDENTITY_PUBLIC_KEY_NAME=WindowsTestTrusted",
               Name, DateTime.Now.ToShortDateString()
               );

            sb.AppendFormat(CultureInfo.InvariantCulture, "\nTARGETNAME      = Microsoft.Test.Wmi.{0}\n", CSNamespaceName);
            sb.AppendLine("SOURCES=\\");

            foreach (KeyValuePair<string, WmiModule> kvpM in Modules)
            {
                foreach (KeyValuePair<string, WmiSource> kvpS in kvpM.Value.Sources)
                {
                    var wSource = kvpS.Value;
                    sb.AppendFormat(CultureInfo.InvariantCulture, "\t{0}\\{1}.cs \\\n", kvpM.Key, wSource.Name);
                }
                foreach (KeyValuePair<string, WmiEnum> kvpS in kvpM.Value.Enums)
                {
                    var wEnum = kvpS.Value;
                    sb.AppendFormat(CultureInfo.InvariantCulture, "\t{0}\\{1}.cs \\\n", kvpM.Key, wEnum.Name);
                }
            }
            sb.AppendLine();
            sb.Append(@"REFERENCES=\
    $(CLR_REF_PATH)\System.metadata_dll; \
    $(CLR_REF_PATH)\System.Collections.metadata_dll; \
    $(CLR_REF_PATH)\System.Core.metadata_dll; \
    $(CLR_REF_PATH)\System.Data.metadata_dll; \
    $(CLR_REF_PATH)\System.Xml.Linq.metadata_dll; \
    $(CLR_REF_PATH)\System.Xml.metadata_dll; \
    $(CLR_REF_PATH)\System.Management.metadata_dll; \
    $(CLR_REF_PATH)\System.ServiceProcess.metadata_dll; \
    $(CLR_REF_PATH)\System.ObjectModel.metadata_dll; \
    $(SDK_REF_PATH)\Microsoft.Management.Infrastructure.metadata_dll; \
    $(OBJECT_ROOT)\net\test\common\wmiwrapper\wmiimplementation\cimwmi\$O\Microsoft.Test.Wmi.Cim.metadata_dll; \
    $(OBJECT_ROOT)\net\test\common\wmiwrapper\wmiinterface\$O\Microsoft.Test.Wmi.metadata_dll; \"
                );

            foreach (var item in sourceReferences)
            {
                sb.AppendLine();
                sb.AppendFormat(CultureInfo.InvariantCulture,
@"   $(OBJECT_ROOT)\net\test\common\wmiwrapper\{0}\{1}\$O\Microsoft.Test.Wmi.{1}.metadata_dll; \", item.Reference, item.Reference.Replace("/", "."));
            }
            File.WriteAllText(sourceFile, sb.ToString());
            Logger.Info("Generated Source file at {0}", sourceFile);
        }

        protected override WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule wModule)
        {
            throw new NotImplementedException();
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }
    }
}