// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: common.proto

package com.github.zhangchengji.flying.common;

/**
 * <pre>
 * 统一响应格式
 * </pre>
 *
 * Protobuf type {@code common.Request}
 */
public final class Request extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:common.Request)
        RequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Request.newBuilder() to construct.
  private Request(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Request() {
    orderKey_ = "";
    orderDesc_ = "";
  }

  @Override
  @SuppressWarnings({"unused"})
  protected Object newInstance(
      UnusedPrivateParameter unused) {
    return new Request();
  }

  @Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Request(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 810: {
            com.google.protobuf.Any.Builder subBuilder = null;
            if (query_ != null) {
              subBuilder = query_.toBuilder();
            }
            query_ = input.readMessage(com.google.protobuf.Any.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(query_);
              query_ = subBuilder.buildPartial();
            }

            break;
          }
          case 816: {

            page_ = input.readInt32();
            break;
          }
          case 824: {

            pageSize_ = input.readInt32();
            break;
          }
          case 834: {
            String s = input.readStringRequireUtf8();

            orderKey_ = s;
            break;
          }
          case 842: {
            String s = input.readStringRequireUtf8();

            orderDesc_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return CommonProto.internal_static_common_Request_descriptor;
  }

  @Override
  protected FieldAccessorTable
      internalGetFieldAccessorTable() {
    return CommonProto.internal_static_common_Request_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            Request.class, Request.Builder.class);
  }

  public static final int QUERY_FIELD_NUMBER = 101;
  private com.google.protobuf.Any query_;
  /**
   * <code>.google.protobuf.Any query = 101;</code>
   * @return Whether the query field is set.
   */
  @Override
  public boolean hasQuery() {
    return query_ != null;
  }
  /**
   * <code>.google.protobuf.Any query = 101;</code>
   * @return The query.
   */
  @Override
  public com.google.protobuf.Any getQuery() {
    return query_ == null ? com.google.protobuf.Any.getDefaultInstance() : query_;
  }
  /**
   * <code>.google.protobuf.Any query = 101;</code>
   */
  @Override
  public com.google.protobuf.AnyOrBuilder getQueryOrBuilder() {
    return getQuery();
  }

  public static final int PAGE_FIELD_NUMBER = 102;
  private int page_;
  /**
   * <code>int32 page = 102;</code>
   * @return The page.
   */
  @Override
  public int getPage() {
    return page_;
  }

  public static final int PAGESIZE_FIELD_NUMBER = 103;
  private int pageSize_;
  /**
   * <code>int32 pageSize = 103;</code>
   * @return The pageSize.
   */
  @Override
  public int getPageSize() {
    return pageSize_;
  }

  public static final int ORDERKEY_FIELD_NUMBER = 104;
  private volatile Object orderKey_;
  /**
   * <code>string orderKey = 104;</code>
   * @return The orderKey.
   */
  @Override
  public String getOrderKey() {
    Object ref = orderKey_;
    if (ref instanceof String) {
      return (String) ref;
    } else {
      com.google.protobuf.ByteString bs =
          (com.google.protobuf.ByteString) ref;
      String s = bs.toStringUtf8();
      orderKey_ = s;
      return s;
    }
  }
  /**
   * <code>string orderKey = 104;</code>
   * @return The bytes for orderKey.
   */
  @Override
  public com.google.protobuf.ByteString
      getOrderKeyBytes() {
    Object ref = orderKey_;
    if (ref instanceof String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8(
              (String) ref);
      orderKey_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int ORDERDESC_FIELD_NUMBER = 105;
  private volatile Object orderDesc_;
  /**
   * <code>string orderDesc = 105;</code>
   * @return The orderDesc.
   */
  @Override
  public String getOrderDesc() {
    Object ref = orderDesc_;
    if (ref instanceof String) {
      return (String) ref;
    } else {
      com.google.protobuf.ByteString bs =
          (com.google.protobuf.ByteString) ref;
      String s = bs.toStringUtf8();
      orderDesc_ = s;
      return s;
    }
  }
  /**
   * <code>string orderDesc = 105;</code>
   * @return The bytes for orderDesc.
   */
  @Override
  public com.google.protobuf.ByteString
      getOrderDescBytes() {
    Object ref = orderDesc_;
    if (ref instanceof String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8(
              (String) ref);
      orderDesc_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (query_ != null) {
      output.writeMessage(101, getQuery());
    }
    if (page_ != 0) {
      output.writeInt32(102, page_);
    }
    if (pageSize_ != 0) {
      output.writeInt32(103, pageSize_);
    }
    if (!getOrderKeyBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 104, orderKey_);
    }
    if (!getOrderDescBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 105, orderDesc_);
    }
    unknownFields.writeTo(output);
  }

  @Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (query_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(101, getQuery());
    }
    if (page_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(102, page_);
    }
    if (pageSize_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(103, pageSize_);
    }
    if (!getOrderKeyBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(104, orderKey_);
    }
    if (!getOrderDescBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(105, orderDesc_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @Override
  public boolean equals(final Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof Request)) {
      return super.equals(obj);
    }
    Request other = (Request) obj;

    if (hasQuery() != other.hasQuery()) return false;
    if (hasQuery()) {
      if (!getQuery()
          .equals(other.getQuery())) return false;
    }
    if (getPage()
        != other.getPage()) return false;
    if (getPageSize()
        != other.getPageSize()) return false;
    if (!getOrderKey()
        .equals(other.getOrderKey())) return false;
    if (!getOrderDesc()
        .equals(other.getOrderDesc())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (hasQuery()) {
      hash = (37 * hash) + QUERY_FIELD_NUMBER;
      hash = (53 * hash) + getQuery().hashCode();
    }
    hash = (37 * hash) + PAGE_FIELD_NUMBER;
    hash = (53 * hash) + getPage();
    hash = (37 * hash) + PAGESIZE_FIELD_NUMBER;
    hash = (53 * hash) + getPageSize();
    hash = (37 * hash) + ORDERKEY_FIELD_NUMBER;
    hash = (53 * hash) + getOrderKey().hashCode();
    hash = (37 * hash) + ORDERDESC_FIELD_NUMBER;
    hash = (53 * hash) + getOrderDesc().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static Request parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static Request parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static Request parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static Request parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static Request parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static Request parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static Request parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static Request parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static Request parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static Request parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static Request parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static Request parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(Request prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @Override
  protected Builder newBuilderForType(
      BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   * 统一响应格式
   * </pre>
   *
   * Protobuf type {@code common.Request}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:common.Request)
          RequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return CommonProto.internal_static_common_Request_descriptor;
    }

    @Override
    protected FieldAccessorTable
        internalGetFieldAccessorTable() {
      return CommonProto.internal_static_common_Request_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              Request.class, Request.Builder.class);
    }

    // Construct using com.github.zhangchengji.proto.common.Request.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @Override
    public Builder clear() {
      super.clear();
      if (queryBuilder_ == null) {
        query_ = null;
      } else {
        query_ = null;
        queryBuilder_ = null;
      }
      page_ = 0;

      pageSize_ = 0;

      orderKey_ = "";

      orderDesc_ = "";

      return this;
    }

    @Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return CommonProto.internal_static_common_Request_descriptor;
    }

    @Override
    public Request getDefaultInstanceForType() {
      return Request.getDefaultInstance();
    }

    @Override
    public Request build() {
      Request result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @Override
    public Request buildPartial() {
      Request result = new Request(this);
      if (queryBuilder_ == null) {
        result.query_ = query_;
      } else {
        result.query_ = queryBuilder_.build();
      }
      result.page_ = page_;
      result.pageSize_ = pageSize_;
      result.orderKey_ = orderKey_;
      result.orderDesc_ = orderDesc_;
      onBuilt();
      return result;
    }

    @Override
    public Builder clone() {
      return super.clone();
    }
    @Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return super.setField(field, value);
    }
    @Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return super.addRepeatedField(field, value);
    }
    @Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof Request) {
        return mergeFrom((Request)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(Request other) {
      if (other == Request.getDefaultInstance()) return this;
      if (other.hasQuery()) {
        mergeQuery(other.getQuery());
      }
      if (other.getPage() != 0) {
        setPage(other.getPage());
      }
      if (other.getPageSize() != 0) {
        setPageSize(other.getPageSize());
      }
      if (!other.getOrderKey().isEmpty()) {
        orderKey_ = other.orderKey_;
        onChanged();
      }
      if (!other.getOrderDesc().isEmpty()) {
        orderDesc_ = other.orderDesc_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @Override
    public final boolean isInitialized() {
      return true;
    }

    @Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      Request parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (Request) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private com.google.protobuf.Any query_;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.google.protobuf.Any, com.google.protobuf.Any.Builder, com.google.protobuf.AnyOrBuilder> queryBuilder_;
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     * @return Whether the query field is set.
     */
    public boolean hasQuery() {
      return queryBuilder_ != null || query_ != null;
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     * @return The query.
     */
    public com.google.protobuf.Any getQuery() {
      if (queryBuilder_ == null) {
        return query_ == null ? com.google.protobuf.Any.getDefaultInstance() : query_;
      } else {
        return queryBuilder_.getMessage();
      }
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    public Builder setQuery(com.google.protobuf.Any value) {
      if (queryBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        query_ = value;
        onChanged();
      } else {
        queryBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    public Builder setQuery(
        com.google.protobuf.Any.Builder builderForValue) {
      if (queryBuilder_ == null) {
        query_ = builderForValue.build();
        onChanged();
      } else {
        queryBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    public Builder mergeQuery(com.google.protobuf.Any value) {
      if (queryBuilder_ == null) {
        if (query_ != null) {
          query_ =
            com.google.protobuf.Any.newBuilder(query_).mergeFrom(value).buildPartial();
        } else {
          query_ = value;
        }
        onChanged();
      } else {
        queryBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    public Builder clearQuery() {
      if (queryBuilder_ == null) {
        query_ = null;
        onChanged();
      } else {
        query_ = null;
        queryBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    public com.google.protobuf.Any.Builder getQueryBuilder() {

      onChanged();
      return getQueryFieldBuilder().getBuilder();
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    public com.google.protobuf.AnyOrBuilder getQueryOrBuilder() {
      if (queryBuilder_ != null) {
        return queryBuilder_.getMessageOrBuilder();
      } else {
        return query_ == null ?
            com.google.protobuf.Any.getDefaultInstance() : query_;
      }
    }
    /**
     * <code>.google.protobuf.Any query = 101;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.google.protobuf.Any, com.google.protobuf.Any.Builder, com.google.protobuf.AnyOrBuilder>
        getQueryFieldBuilder() {
      if (queryBuilder_ == null) {
        queryBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.google.protobuf.Any, com.google.protobuf.Any.Builder, com.google.protobuf.AnyOrBuilder>(
                getQuery(),
                getParentForChildren(),
                isClean());
        query_ = null;
      }
      return queryBuilder_;
    }

    private int page_ ;
    /**
     * <code>int32 page = 102;</code>
     * @return The page.
     */
    @Override
    public int getPage() {
      return page_;
    }
    /**
     * <code>int32 page = 102;</code>
     * @param value The page to set.
     * @return This builder for chaining.
     */
    public Builder setPage(int value) {

      page_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 page = 102;</code>
     * @return This builder for chaining.
     */
    public Builder clearPage() {

      page_ = 0;
      onChanged();
      return this;
    }

    private int pageSize_ ;
    /**
     * <code>int32 pageSize = 103;</code>
     * @return The pageSize.
     */
    @Override
    public int getPageSize() {
      return pageSize_;
    }
    /**
     * <code>int32 pageSize = 103;</code>
     * @param value The pageSize to set.
     * @return This builder for chaining.
     */
    public Builder setPageSize(int value) {

      pageSize_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 pageSize = 103;</code>
     * @return This builder for chaining.
     */
    public Builder clearPageSize() {

      pageSize_ = 0;
      onChanged();
      return this;
    }

    private Object orderKey_ = "";
    /**
     * <code>string orderKey = 104;</code>
     * @return The orderKey.
     */
    public String getOrderKey() {
      Object ref = orderKey_;
      if (!(ref instanceof String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        orderKey_ = s;
        return s;
      } else {
        return (String) ref;
      }
    }
    /**
     * <code>string orderKey = 104;</code>
     * @return The bytes for orderKey.
     */
    public com.google.protobuf.ByteString
        getOrderKeyBytes() {
      Object ref = orderKey_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        orderKey_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string orderKey = 104;</code>
     * @param value The orderKey to set.
     * @return This builder for chaining.
     */
    public Builder setOrderKey(
        String value) {
      if (value == null) {
    throw new NullPointerException();
  }

      orderKey_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string orderKey = 104;</code>
     * @return This builder for chaining.
     */
    public Builder clearOrderKey() {

      orderKey_ = getDefaultInstance().getOrderKey();
      onChanged();
      return this;
    }
    /**
     * <code>string orderKey = 104;</code>
     * @param value The bytes for orderKey to set.
     * @return This builder for chaining.
     */
    public Builder setOrderKeyBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);

      orderKey_ = value;
      onChanged();
      return this;
    }

    private Object orderDesc_ = "";
    /**
     * <code>string orderDesc = 105;</code>
     * @return The orderDesc.
     */
    public String getOrderDesc() {
      Object ref = orderDesc_;
      if (!(ref instanceof String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        orderDesc_ = s;
        return s;
      } else {
        return (String) ref;
      }
    }
    /**
     * <code>string orderDesc = 105;</code>
     * @return The bytes for orderDesc.
     */
    public com.google.protobuf.ByteString
        getOrderDescBytes() {
      Object ref = orderDesc_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        orderDesc_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string orderDesc = 105;</code>
     * @param value The orderDesc to set.
     * @return This builder for chaining.
     */
    public Builder setOrderDesc(
        String value) {
      if (value == null) {
    throw new NullPointerException();
  }

      orderDesc_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string orderDesc = 105;</code>
     * @return This builder for chaining.
     */
    public Builder clearOrderDesc() {

      orderDesc_ = getDefaultInstance().getOrderDesc();
      onChanged();
      return this;
    }
    /**
     * <code>string orderDesc = 105;</code>
     * @param value The bytes for orderDesc to set.
     * @return This builder for chaining.
     */
    public Builder setOrderDescBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);

      orderDesc_ = value;
      onChanged();
      return this;
    }
    @Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:common.Request)
  }

  // @@protoc_insertion_point(class_scope:common.Request)
  private static final Request DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new Request();
  }

  public static Request getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Request>
      PARSER = new com.google.protobuf.AbstractParser<Request>() {
    @Override
    public Request parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Request(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Request> parser() {
    return PARSER;
  }

  @Override
  public com.google.protobuf.Parser<Request> getParserForType() {
    return PARSER;
  }

  @Override
  public Request getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

