/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.Set;
import java.util.HashSet;
import java.util.Collections;
import java.util.BitSet;
import java.util.Arrays;
import com.facebook.thrift.*;
import com.facebook.thrift.annotations.*;
import com.facebook.thrift.async.*;
import com.facebook.thrift.meta_data.*;
import com.facebook.thrift.server.*;
import com.facebook.thrift.transport.*;
import com.facebook.thrift.protocol.*;

@SuppressWarnings({ "unused", "serial" })
public class AdaptedStructWithInternBox implements TBase, java.io.Serializable, Cloneable {
  private static final TStruct STRUCT_DESC = new TStruct("AdaptedStructWithInternBox");
  private static final TField FIELD1_FIELD_DESC = new TField("field1", TType.STRUCT, (short)1);
  private static final TField FIELD2_FIELD_DESC = new TField("field2", TType.STRUCT, (short)2);

  public final Empty field1;
  public final MyField field2;
  public static final int FIELD1 = 1;
  public static final int FIELD2 = 2;

  public AdaptedStructWithInternBox(
      Empty field1,
      MyField field2) {
    this.field1 = field1;
    this.field2 = field2;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public AdaptedStructWithInternBox(AdaptedStructWithInternBox other) {
    if (other.isSetField1()) {
      this.field1 = TBaseHelper.deepCopy(other.field1);
    } else {
      this.field1 = null;
    }
    if (other.isSetField2()) {
      this.field2 = TBaseHelper.deepCopy(other.field2);
    } else {
      this.field2 = null;
    }
  }

  public AdaptedStructWithInternBox deepCopy() {
    return new AdaptedStructWithInternBox(this);
  }

  public Empty getField1() {
    return this.field1;
  }

  // Returns true if field field1 is set (has been assigned a value) and false otherwise
  public boolean isSetField1() {
    return this.field1 != null;
  }

  public MyField getField2() {
    return this.field2;
  }

  // Returns true if field field2 is set (has been assigned a value) and false otherwise
  public boolean isSetField2() {
    return this.field2 != null;
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof AdaptedStructWithInternBox))
      return false;
    AdaptedStructWithInternBox that = (AdaptedStructWithInternBox)_that;

    if (!TBaseHelper.equalsNobinary(this.isSetField1(), that.isSetField1(), this.field1, that.field1)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetField2(), that.isSetField2(), this.field2, that.field2)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {field1, field2});
  }

  // This is required to satisfy the TBase interface, but can't be implemented on immutable struture.
  public void read(TProtocol iprot) throws TException {
    throw new TException("unimplemented in android immutable structure");
  }

  public static AdaptedStructWithInternBox deserialize(TProtocol iprot) throws TException {
    Empty tmp_field1 = null;
    MyField tmp_field2 = null;
    TField __field;
    iprot.readStructBegin();
    while (true)
    {
      __field = iprot.readFieldBegin();
      if (__field.type == TType.STOP) {
        break;
      }
      switch (__field.id)
      {
        case FIELD1:
          if (__field.type == TType.STRUCT) {
            tmp_field1 = Empty.deserialize(iprot);
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case FIELD2:
          if (__field.type == TType.STRUCT) {
            tmp_field2 = MyField.deserialize(iprot);
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        default:
          TProtocolUtil.skip(iprot, __field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    AdaptedStructWithInternBox _that;
    _that = new AdaptedStructWithInternBox(
      tmp_field1
      ,tmp_field2
    );
    _that.validate();
    return _that;
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    if (this.field1 != null) {
      oprot.writeFieldBegin(FIELD1_FIELD_DESC);
      this.field1.write(oprot);
      oprot.writeFieldEnd();
    }
    if (this.field2 != null) {
      oprot.writeFieldBegin(FIELD2_FIELD_DESC);
      this.field2.write(oprot);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @Override
  public String toString() {
    return toString(1, true);
  }

  @Override
  public String toString(int indent, boolean prettyPrint) {
    return TBaseHelper.toStringHelper(this, indent, prettyPrint);
  }

  public void validate() throws TException {
    // check for required fields
  }

}
