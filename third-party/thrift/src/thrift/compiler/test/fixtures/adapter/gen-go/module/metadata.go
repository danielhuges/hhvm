// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package module // [[[ program thrift source path ]]]

import (
    "maps"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = maps.Copy[map[int]int, map[int]int]
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
            )
    premadeThriftType_module_Color = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("module.Color"),
            )
    premadeThriftType_i32 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I32_TYPE.Ptr(),
            )
    premadeThriftType_module_i32_5137 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.i32_5137").
            SetUnderlyingType(premadeThriftType_i32),
            )
    premadeThriftType_set_string = metadata.NewThriftType().SetTSet(
        metadata.NewThriftSetType().
            SetValueType(premadeThriftType_string),
            )
    premadeThriftType_module_SetWithAdapter = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.SetWithAdapter").
            SetUnderlyingType(premadeThriftType_set_string),
            )
    premadeThriftType_module_StringWithAdapter = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.StringWithAdapter").
            SetUnderlyingType(premadeThriftType_string),
            )
    premadeThriftType_list_module_StringWithAdapter = metadata.NewThriftType().SetTList(
        metadata.NewThriftListType().
            SetValueType(premadeThriftType_module_StringWithAdapter),
            )
    premadeThriftType_module_ListWithElemAdapter = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.ListWithElemAdapter").
            SetUnderlyingType(premadeThriftType_list_module_StringWithAdapter),
            )
    premadeThriftType_module_ListWithElemAdapter_withAdapter = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.ListWithElemAdapter_withAdapter").
            SetUnderlyingType(premadeThriftType_module_ListWithElemAdapter),
            )
    premadeThriftType_module_ListWithElemAdapter_withAdapter_2312 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.ListWithElemAdapter_withAdapter_2312").
            SetUnderlyingType(premadeThriftType_module_ListWithElemAdapter_withAdapter),
            )
    premadeThriftType_map_string_module_ListWithElemAdapter_withAdapter_2312 = metadata.NewThriftType().SetTMap(
        metadata.NewThriftMapType().
            SetKeyType(premadeThriftType_string).
            SetValueType(premadeThriftType_module_ListWithElemAdapter_withAdapter_2312),
            )
    premadeThriftType_module_map_string_ListWithElemAdapter_withAdapter_8454 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.map_string_ListWithElemAdapter_withAdapter_8454").
            SetUnderlyingType(premadeThriftType_map_string_module_ListWithElemAdapter_withAdapter_2312),
            )
    premadeThriftType_binary = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BINARY_TYPE.Ptr(),
            )
    premadeThriftType_module_binary_5673 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.binary_5673").
            SetUnderlyingType(premadeThriftType_binary),
            )
    premadeThriftType_i64 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I64_TYPE.Ptr(),
            )
    premadeThriftType_module_MyI64 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.MyI64").
            SetUnderlyingType(premadeThriftType_i64),
            )
    premadeThriftType_module_DoubleTypedefI64 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.DoubleTypedefI64").
            SetUnderlyingType(premadeThriftType_module_MyI64),
            )
    premadeThriftType_module_Foo = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.Foo"),
            )
    premadeThriftType_module_Foo_6868 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.Foo_6868").
            SetUnderlyingType(premadeThriftType_module_Foo),
            )
    premadeThriftType_module_Foo_3943 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.Foo_3943").
            SetUnderlyingType(premadeThriftType_module_Foo),
            )
    premadeThriftType_module_FooWithAdapter = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.FooWithAdapter").
            SetUnderlyingType(premadeThriftType_module_Foo),
            )
    premadeThriftType_module_FooWithAdapter_9317 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.FooWithAdapter_9317").
            SetUnderlyingType(premadeThriftType_module_FooWithAdapter),
            )
    premadeThriftType_list_module_FooWithAdapter_9317 = metadata.NewThriftType().SetTList(
        metadata.NewThriftListType().
            SetValueType(premadeThriftType_module_FooWithAdapter_9317),
            )
    premadeThriftType_module_Baz = metadata.NewThriftType().SetTUnion(
        metadata.NewThriftUnionType().
            SetName("module.Baz"),
            )
    premadeThriftType_module_Baz_7352 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.Baz_7352").
            SetUnderlyingType(premadeThriftType_module_Baz),
            )
    premadeThriftType_module_DirectlyAdapted = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.DirectlyAdapted"),
            )
    premadeThriftType_set_i32 = metadata.NewThriftType().SetTSet(
        metadata.NewThriftSetType().
            SetValueType(premadeThriftType_i32),
            )
    premadeThriftType_module_A = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.A"),
            )
    premadeThriftType_module_AdaptedA = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedA").
            SetUnderlyingType(premadeThriftType_module_A),
            )
    premadeThriftType_module_DurationMs = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.DurationMs").
            SetUnderlyingType(premadeThriftType_i64),
            )
    premadeThriftType_module_IOBuf = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.IOBuf").
            SetUnderlyingType(premadeThriftType_binary),
            )
    premadeThriftType_module_CustomProtocolType = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.CustomProtocolType").
            SetUnderlyingType(premadeThriftType_module_IOBuf),
            )
    premadeThriftType_module_IndirectionString = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.IndirectionString").
            SetUnderlyingType(premadeThriftType_string),
            )
    premadeThriftType_bool = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BOOL_TYPE.Ptr(),
            )
    premadeThriftType_module_AdaptedBool = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedBool").
            SetUnderlyingType(premadeThriftType_bool),
            )
    premadeThriftType_module_AdaptedInteger = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedInteger").
            SetUnderlyingType(premadeThriftType_i32),
            )
    premadeThriftType_byte = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BYTE_TYPE.Ptr(),
            )
    premadeThriftType_module_AdaptedByte = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedByte").
            SetUnderlyingType(premadeThriftType_byte),
            )
    premadeThriftType_i16 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I16_TYPE.Ptr(),
            )
    premadeThriftType_module_AdaptedShort = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedShort").
            SetUnderlyingType(premadeThriftType_i16),
            )
    premadeThriftType_module_AdaptedLong = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedLong").
            SetUnderlyingType(premadeThriftType_i64),
            )
    premadeThriftType_double = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_DOUBLE_TYPE.Ptr(),
            )
    premadeThriftType_module_AdaptedDouble = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedDouble").
            SetUnderlyingType(premadeThriftType_double),
            )
    premadeThriftType_module_AdaptedString = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedString").
            SetUnderlyingType(premadeThriftType_string),
            )
    premadeThriftType_list_i64 = metadata.NewThriftType().SetTList(
        metadata.NewThriftListType().
            SetValueType(premadeThriftType_i64),
            )
    premadeThriftType_set_i64 = metadata.NewThriftType().SetTSet(
        metadata.NewThriftSetType().
            SetValueType(premadeThriftType_i64),
            )
    premadeThriftType_map_i64_i64 = metadata.NewThriftType().SetTMap(
        metadata.NewThriftMapType().
            SetKeyType(premadeThriftType_i64).
            SetValueType(premadeThriftType_i64),
            )
    premadeThriftType_module_ThriftAdaptedEnum = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("module.ThriftAdaptedEnum"),
            )
    premadeThriftType_module_AdaptedEnum = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedEnum").
            SetUnderlyingType(premadeThriftType_module_ThriftAdaptedEnum),
            )
    premadeThriftType_module_DoubleTypedefBool = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.DoubleTypedefBool").
            SetUnderlyingType(premadeThriftType_module_AdaptedBool),
            )
    premadeThriftType_module_AdaptTemplatedTestStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.AdaptTemplatedTestStruct"),
            )
    premadeThriftType_module_AdaptedStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.AdaptedStruct"),
            )
    premadeThriftType_module_AdaptedTypedef = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedTypedef").
            SetUnderlyingType(premadeThriftType_module_AdaptedStruct),
            )
    premadeThriftType_module_DirectlyAdaptedStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.DirectlyAdaptedStruct"),
            )
    premadeThriftType_module_TypedefOfDirect = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.TypedefOfDirect").
            SetUnderlyingType(premadeThriftType_module_DirectlyAdaptedStruct),
            )
    premadeThriftType_module_CircularStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.CircularStruct"),
            )
    premadeThriftType_module_CircularAdaptee = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.CircularAdaptee"),
            )
    premadeThriftType_module_AdaptedCircularAdaptee = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.AdaptedCircularAdaptee").
            SetUnderlyingType(premadeThriftType_module_CircularAdaptee),
            )
    premadeThriftType_module_DeclaredAfterStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.DeclaredAfterStruct"),
            )
    premadeThriftType_module_HeapAllocated = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.HeapAllocated"),
            )
    premadeThriftType_module_CountingInt = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.CountingInt").
            SetUnderlyingType(premadeThriftType_i64),
            )
    premadeThriftType_module_Bar = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.Bar"),
            )
    premadeThriftType_module_MyI32 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.MyI32").
            SetUnderlyingType(premadeThriftType_i32),
            )
    premadeThriftType_module_MyI32_4873 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.MyI32_4873").
            SetUnderlyingType(premadeThriftType_module_MyI32),
            )
    premadeThriftType_module_StringWithAdapter_7208 = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("module.StringWithAdapter_7208").
            SetUnderlyingType(premadeThriftType_module_StringWithAdapter),
            )
    premadeThriftType_module_CountingStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.CountingStruct"),
            )
)

var structMetadatas = []*metadata.ThriftStruct{
    metadata.NewThriftStruct().
    SetName("module.MyAnnotation").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("signature").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(2).
    SetName("color").
    SetIsOptional(false).
    SetType(premadeThriftType_module_Color),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.Foo").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("intField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_i32_5137),
            metadata.NewThriftField().
    SetId(2).
    SetName("optionalIntField").
    SetIsOptional(true).
    SetType(premadeThriftType_module_i32_5137),
            metadata.NewThriftField().
    SetId(3).
    SetName("intFieldWithDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_i32_5137),
            metadata.NewThriftField().
    SetId(4).
    SetName("setField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_SetWithAdapter),
            metadata.NewThriftField().
    SetId(5).
    SetName("optionalSetField").
    SetIsOptional(true).
    SetType(premadeThriftType_module_SetWithAdapter),
            metadata.NewThriftField().
    SetId(6).
    SetName("mapField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_map_string_ListWithElemAdapter_withAdapter_8454),
            metadata.NewThriftField().
    SetId(7).
    SetName("optionalMapField").
    SetIsOptional(true).
    SetType(premadeThriftType_module_map_string_ListWithElemAdapter_withAdapter_8454),
            metadata.NewThriftField().
    SetId(8).
    SetName("binaryField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_binary_5673),
            metadata.NewThriftField().
    SetId(9).
    SetName("longField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyI64),
            metadata.NewThriftField().
    SetId(10).
    SetName("adaptedLongField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyI64),
            metadata.NewThriftField().
    SetId(11).
    SetName("doubleAdaptedField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DoubleTypedefI64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.Baz").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("intField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_i32_5137),
            metadata.NewThriftField().
    SetId(4).
    SetName("setField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_SetWithAdapter),
            metadata.NewThriftField().
    SetId(6).
    SetName("mapField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_map_string_ListWithElemAdapter_withAdapter_8454),
            metadata.NewThriftField().
    SetId(8).
    SetName("binaryField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_binary_5673),
            metadata.NewThriftField().
    SetId(9).
    SetName("longField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyI64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.Bar").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("structField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_Foo_6868),
            metadata.NewThriftField().
    SetId(2).
    SetName("optionalStructField").
    SetIsOptional(true).
    SetType(premadeThriftType_module_Foo_3943),
            metadata.NewThriftField().
    SetId(3).
    SetName("structListField").
    SetIsOptional(false).
    SetType(premadeThriftType_list_module_FooWithAdapter_9317),
            metadata.NewThriftField().
    SetId(4).
    SetName("optionalStructListField").
    SetIsOptional(true).
    SetType(premadeThriftType_list_module_FooWithAdapter_9317),
            metadata.NewThriftField().
    SetId(5).
    SetName("unionField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_Baz_7352),
            metadata.NewThriftField().
    SetId(6).
    SetName("optionalUnionField").
    SetIsOptional(true).
    SetType(premadeThriftType_module_Baz_7352),
            metadata.NewThriftField().
    SetId(7).
    SetName("adaptedStructField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DirectlyAdapted),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.DirectlyAdapted").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.IndependentDirectlyAdapted").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.StructWithFieldAdapter").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(2).
    SetName("shared_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(3).
    SetName("opt_shared_field").
    SetIsOptional(true).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(4).
    SetName("opt_boxed_field").
    SetIsOptional(true).
    SetType(premadeThriftType_i32),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.TerseAdaptedFields").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("int_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(2).
    SetName("string_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(3).
    SetName("set_field").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i32),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.B").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("a").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedA),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.A").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("module.Config").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("path").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.MyStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(2).
    SetName("set_string").
    SetIsOptional(false).
    SetType(premadeThriftType_module_SetWithAdapter),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.AdaptTestStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("delay").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DurationMs),
            metadata.NewThriftField().
    SetId(2).
    SetName("custom").
    SetIsOptional(false).
    SetType(premadeThriftType_module_CustomProtocolType),
            metadata.NewThriftField().
    SetId(3).
    SetName("timeout").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(4).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(5).
    SetName("meta").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(6).
    SetName("indirectionString").
    SetIsOptional(false).
    SetType(premadeThriftType_module_IndirectionString),
            metadata.NewThriftField().
    SetId(7).
    SetName("string_data").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(8).
    SetName("double_wrapped_bool").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedBool),
            metadata.NewThriftField().
    SetId(9).
    SetName("double_wrapped_integer").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedInteger),
            metadata.NewThriftField().
    SetId(10).
    SetName("binary_data").
    SetIsOptional(false).
    SetType(premadeThriftType_binary),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.AdaptTemplatedTestStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("adaptedBool").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedBool),
            metadata.NewThriftField().
    SetId(2).
    SetName("adaptedByte").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedByte),
            metadata.NewThriftField().
    SetId(3).
    SetName("adaptedShort").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedShort),
            metadata.NewThriftField().
    SetId(4).
    SetName("adaptedInteger").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedInteger),
            metadata.NewThriftField().
    SetId(5).
    SetName("adaptedLong").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedLong),
            metadata.NewThriftField().
    SetId(6).
    SetName("adaptedDouble").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedDouble),
            metadata.NewThriftField().
    SetId(7).
    SetName("adaptedString").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedString),
            metadata.NewThriftField().
    SetId(8).
    SetName("adaptedList").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i64),
            metadata.NewThriftField().
    SetId(9).
    SetName("adaptedSet").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i64),
            metadata.NewThriftField().
    SetId(10).
    SetName("adaptedMap").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i64_i64),
            metadata.NewThriftField().
    SetId(11).
    SetName("adaptedBoolDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedBool),
            metadata.NewThriftField().
    SetId(12).
    SetName("adaptedByteDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedByte),
            metadata.NewThriftField().
    SetId(13).
    SetName("adaptedShortDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedShort),
            metadata.NewThriftField().
    SetId(14).
    SetName("adaptedIntegerDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedInteger),
            metadata.NewThriftField().
    SetId(15).
    SetName("adaptedLongDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedLong),
            metadata.NewThriftField().
    SetId(16).
    SetName("adaptedDoubleDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedDouble),
            metadata.NewThriftField().
    SetId(17).
    SetName("adaptedStringDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedString),
            metadata.NewThriftField().
    SetId(18).
    SetName("adaptedEnum").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedEnum),
            metadata.NewThriftField().
    SetId(19).
    SetName("adaptedListDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i64),
            metadata.NewThriftField().
    SetId(20).
    SetName("adaptedSetDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i64),
            metadata.NewThriftField().
    SetId(21).
    SetName("adaptedMapDefault").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i64_i64),
            metadata.NewThriftField().
    SetId(22).
    SetName("doubleTypedefBool").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DoubleTypedefBool),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.AdaptTemplatedNestedTestStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("adaptedStruct").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptTemplatedTestStruct),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.AdaptTestUnion").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("delay").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DurationMs),
            metadata.NewThriftField().
    SetId(2).
    SetName("custom").
    SetIsOptional(false).
    SetType(premadeThriftType_module_CustomProtocolType),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.AdaptedStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.DirectlyAdaptedStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.StructFieldAdaptedStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("adaptedStruct").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedStruct),
            metadata.NewThriftField().
    SetId(2).
    SetName("adaptedTypedef").
    SetIsOptional(false).
    SetType(premadeThriftType_module_AdaptedTypedef),
            metadata.NewThriftField().
    SetId(3).
    SetName("directlyAdapted").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DirectlyAdaptedStruct),
            metadata.NewThriftField().
    SetId(4).
    SetName("typedefOfAdapted").
    SetIsOptional(false).
    SetType(premadeThriftType_module_TypedefOfDirect),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.CircularAdaptee").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field").
    SetIsOptional(false).
    SetType(premadeThriftType_module_CircularStruct),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.CircularStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field").
    SetIsOptional(true).
    SetType(premadeThriftType_module_AdaptedCircularAdaptee),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.ReorderedStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("reordered_dependent_adapted").
    SetIsOptional(false).
    SetType(premadeThriftType_module_DeclaredAfterStruct),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.DeclaredAfterStruct").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("module.RenamedStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.SameNamespaceStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.HeapAllocated").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("module.MoveOnly").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("ptr").
    SetIsOptional(false).
    SetType(premadeThriftType_module_HeapAllocated),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.AlsoMoveOnly").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("ptr").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.ApplyAdapter").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("module.TransitiveAdapted").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("module.CountingStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("regularInt").
    SetIsOptional(true).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(2).
    SetName("countingInt").
    SetIsOptional(true).
    SetType(premadeThriftType_module_CountingInt),
            metadata.NewThriftField().
    SetId(3).
    SetName("regularString").
    SetIsOptional(true).
    SetType(premadeThriftType_string),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.Person").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("name").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.Person2").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("name").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
}

var exceptionMetadatas = []*metadata.ThriftException{
}

var enumMetadatas = []*metadata.ThriftEnum{
    metadata.NewThriftEnum().
    SetName("module.Color").
    SetElements(
        map[int32]string{
            0: "UNKNOWN",
            1: "RED",
            2: "GREEN",
            3: "BLUE",
        },
    ),
    metadata.NewThriftEnum().
    SetName("module.ThriftAdaptedEnum").
    SetElements(
        map[int32]string{
            0: "Zero",
            1: "One",
        },
    ),
}

var serviceMetadatas = []*metadata.ThriftService{
    metadata.NewThriftService().
    SetName("module.Service").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("func").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_module_MyI32_4873).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("arg1").
    SetIsOptional(false).
    SetType(premadeThriftType_module_StringWithAdapter_7208),
            metadata.NewThriftField().
    SetId(2).
    SetName("arg2").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(3).
    SetName("arg3").
    SetIsOptional(false).
    SetType(premadeThriftType_module_Foo),
        },
    ),
        },
    ),
    metadata.NewThriftService().
    SetName("module.AdapterService").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("count").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_module_CountingStruct),
            metadata.NewThriftFunction().
    SetName("adaptedTypes").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_module_HeapAllocated).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("arg").
    SetIsOptional(false).
    SetType(premadeThriftType_module_HeapAllocated),
        },
    ),
        },
    ),
}

// GetThriftMetadata returns complete Thrift metadata for current and imported packages.
func GetThriftMetadata() *metadata.ThriftMetadata {
    allEnums := GetEnumsMetadata()
    allStructs := GetStructsMetadata()
    allExceptions := GetExceptionsMetadata()
    allServices := GetServicesMetadata()

    return metadata.NewThriftMetadata().
        SetEnums(allEnums).
        SetStructs(allStructs).
        SetExceptions(allExceptions).
        SetServices(allServices)
}

// GetEnumsMetadata returns Thrift metadata for enums in the current and recursively included packages.
func GetEnumsMetadata() map[string]*metadata.ThriftEnum {
    allEnumsMap := make(map[string]*metadata.ThriftEnum)

    // Add enum metadatas from the current program...
    for _, enumMetadata := range enumMetadatas {
        allEnumsMap[enumMetadata.GetName()] = enumMetadata
    }

    // ...now add enum metadatas from recursively included programs.

    return allEnumsMap
}

// GetStructsMetadata returns Thrift metadata for structs in the current and recursively included packages.
func GetStructsMetadata() map[string]*metadata.ThriftStruct {
    allStructsMap := make(map[string]*metadata.ThriftStruct)

    // Add struct metadatas from the current program...
    for _, structMetadata := range structMetadatas {
        allStructsMap[structMetadata.GetName()] = structMetadata
    }

    // ...now add struct metadatas from recursively included programs.

    return allStructsMap
}

// GetExceptionsMetadata returns Thrift metadata for exceptions in the current and recursively included packages.
func GetExceptionsMetadata() map[string]*metadata.ThriftException {
    allExceptionsMap := make(map[string]*metadata.ThriftException)

    // Add exception metadatas from the current program...
    for _, exceptionMetadata := range exceptionMetadatas {
        allExceptionsMap[exceptionMetadata.GetName()] = exceptionMetadata
    }

    // ...now add exception metadatas from recursively included programs.

    return allExceptionsMap
}

// GetServicesMetadata returns Thrift metadata for services in the current and recursively included packages.
func GetServicesMetadata() map[string]*metadata.ThriftService {
    allServicesMap := make(map[string]*metadata.ThriftService)

    // Add service metadatas from the current program...
    for _, serviceMetadata := range serviceMetadatas {
        allServicesMap[serviceMetadata.GetName()] = serviceMetadata
    }

    // ...now add service metadatas from recursively included programs.

    return allServicesMap
}

// GetThriftMetadataForService returns Thrift metadata for the given service.
func GetThriftMetadataForService(scopedServiceName string) *metadata.ThriftMetadata {
    thriftMetadata := GetThriftMetadata()

    allServicesMap := thriftMetadata.GetServices()
    relevantServicesMap := make(map[string]*metadata.ThriftService)

    serviceMetadata := allServicesMap[scopedServiceName]
    // Visit and record all recursive parents of the target service.
    for serviceMetadata != nil {
        relevantServicesMap[serviceMetadata.GetName()] = serviceMetadata
        if serviceMetadata.IsSetParent() {
            serviceMetadata = allServicesMap[serviceMetadata.GetParent()]
        } else {
            serviceMetadata = nil
        }
    }

    thriftMetadata.SetServices(relevantServicesMap)

    return thriftMetadata
}