// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: common.proto

package com.github.zhangchengji.flying.common;

public interface ResponsesOrBuilder extends
    // @@protoc_insertion_point(interface_extends:common.Responses)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int32 code = 201;</code>
   * @return The code.
   */
  int getCode();

  /**
   * <code>string message = 202;</code>
   * @return The message.
   */
  String getMessage();
  /**
   * <code>string message = 202;</code>
   * @return The bytes for message.
   */
  com.google.protobuf.ByteString
      getMessageBytes();

  /**
   * <code>repeated .google.protobuf.Any data = 203;</code>
   */
  java.util.List<com.google.protobuf.Any> 
      getDataList();
  /**
   * <code>repeated .google.protobuf.Any data = 203;</code>
   */
  com.google.protobuf.Any getData(int index);
  /**
   * <code>repeated .google.protobuf.Any data = 203;</code>
   */
  int getDataCount();
  /**
   * <code>repeated .google.protobuf.Any data = 203;</code>
   */
  java.util.List<? extends com.google.protobuf.AnyOrBuilder> 
      getDataOrBuilderList();
  /**
   * <code>repeated .google.protobuf.Any data = 203;</code>
   */
  com.google.protobuf.AnyOrBuilder getDataOrBuilder(
      int index);

  /**
   * <code>.common.Meta meta = 204;</code>
   * @return Whether the meta field is set.
   */
  boolean hasMeta();
  /**
   * <code>.common.Meta meta = 204;</code>
   * @return The meta.
   */
  Meta getMeta();
  /**
   * <code>.common.Meta meta = 204;</code>
   */
  MetaOrBuilder getMetaOrBuilder();
}
