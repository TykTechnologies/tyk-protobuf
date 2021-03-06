// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: coprocess_return_overrides.proto

#ifndef PROTOBUF_INCLUDED_coprocess_5freturn_5foverrides_2eproto
#define PROTOBUF_INCLUDED_coprocess_5freturn_5foverrides_2eproto

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3006001
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3006001 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/map.h>  // IWYU pragma: export
#include <google/protobuf/map_entry.h>
#include <google/protobuf/map_field_inl.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#define PROTOBUF_INTERNAL_EXPORT_protobuf_coprocess_5freturn_5foverrides_2eproto 

namespace protobuf_coprocess_5freturn_5foverrides_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[2];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
}  // namespace protobuf_coprocess_5freturn_5foverrides_2eproto
namespace coprocess {
class ReturnOverrides;
class ReturnOverridesDefaultTypeInternal;
extern ReturnOverridesDefaultTypeInternal _ReturnOverrides_default_instance_;
class ReturnOverrides_HeadersEntry_DoNotUse;
class ReturnOverrides_HeadersEntry_DoNotUseDefaultTypeInternal;
extern ReturnOverrides_HeadersEntry_DoNotUseDefaultTypeInternal _ReturnOverrides_HeadersEntry_DoNotUse_default_instance_;
}  // namespace coprocess
namespace google {
namespace protobuf {
template<> ::coprocess::ReturnOverrides* Arena::CreateMaybeMessage<::coprocess::ReturnOverrides>(Arena*);
template<> ::coprocess::ReturnOverrides_HeadersEntry_DoNotUse* Arena::CreateMaybeMessage<::coprocess::ReturnOverrides_HeadersEntry_DoNotUse>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace coprocess {

// ===================================================================

class ReturnOverrides_HeadersEntry_DoNotUse : public ::google::protobuf::internal::MapEntry<ReturnOverrides_HeadersEntry_DoNotUse, 
    ::std::string, ::std::string,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    0 > {
public:
  typedef ::google::protobuf::internal::MapEntry<ReturnOverrides_HeadersEntry_DoNotUse, 
    ::std::string, ::std::string,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    0 > SuperType;
  ReturnOverrides_HeadersEntry_DoNotUse();
  ReturnOverrides_HeadersEntry_DoNotUse(::google::protobuf::Arena* arena);
  void MergeFrom(const ReturnOverrides_HeadersEntry_DoNotUse& other);
  static const ReturnOverrides_HeadersEntry_DoNotUse* internal_default_instance() { return reinterpret_cast<const ReturnOverrides_HeadersEntry_DoNotUse*>(&_ReturnOverrides_HeadersEntry_DoNotUse_default_instance_); }
  void MergeFrom(const ::google::protobuf::Message& other) final;
  ::google::protobuf::Metadata GetMetadata() const;
};

// -------------------------------------------------------------------

class ReturnOverrides : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:coprocess.ReturnOverrides) */ {
 public:
  ReturnOverrides();
  virtual ~ReturnOverrides();

  ReturnOverrides(const ReturnOverrides& from);

  inline ReturnOverrides& operator=(const ReturnOverrides& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  ReturnOverrides(ReturnOverrides&& from) noexcept
    : ReturnOverrides() {
    *this = ::std::move(from);
  }

  inline ReturnOverrides& operator=(ReturnOverrides&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const ReturnOverrides& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const ReturnOverrides* internal_default_instance() {
    return reinterpret_cast<const ReturnOverrides*>(
               &_ReturnOverrides_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(ReturnOverrides* other);
  friend void swap(ReturnOverrides& a, ReturnOverrides& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline ReturnOverrides* New() const final {
    return CreateMaybeMessage<ReturnOverrides>(NULL);
  }

  ReturnOverrides* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<ReturnOverrides>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const ReturnOverrides& from);
  void MergeFrom(const ReturnOverrides& from);
  void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(ReturnOverrides* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------


  // accessors -------------------------------------------------------

  // map<string, string> headers = 3;
  int headers_size() const;
  void clear_headers();
  static const int kHeadersFieldNumber = 3;
  const ::google::protobuf::Map< ::std::string, ::std::string >&
      headers() const;
  ::google::protobuf::Map< ::std::string, ::std::string >*
      mutable_headers();

  // string response_error = 2;
  void clear_response_error();
  static const int kResponseErrorFieldNumber = 2;
  const ::std::string& response_error() const;
  void set_response_error(const ::std::string& value);
  #if LANG_CXX11
  void set_response_error(::std::string&& value);
  #endif
  void set_response_error(const char* value);
  void set_response_error(const char* value, size_t size);
  ::std::string* mutable_response_error();
  ::std::string* release_response_error();
  void set_allocated_response_error(::std::string* response_error);

  // int32 response_code = 1;
  void clear_response_code();
  static const int kResponseCodeFieldNumber = 1;
  ::google::protobuf::int32 response_code() const;
  void set_response_code(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:coprocess.ReturnOverrides)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::MapField<
      ReturnOverrides_HeadersEntry_DoNotUse,
      ::std::string, ::std::string,
      ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
      ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
      0 > headers_;
  ::google::protobuf::internal::ArenaStringPtr response_error_;
  ::google::protobuf::int32 response_code_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_coprocess_5freturn_5foverrides_2eproto::TableStruct;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// ReturnOverrides

// int32 response_code = 1;
inline void ReturnOverrides::clear_response_code() {
  response_code_ = 0;
}
inline ::google::protobuf::int32 ReturnOverrides::response_code() const {
  // @@protoc_insertion_point(field_get:coprocess.ReturnOverrides.response_code)
  return response_code_;
}
inline void ReturnOverrides::set_response_code(::google::protobuf::int32 value) {
  
  response_code_ = value;
  // @@protoc_insertion_point(field_set:coprocess.ReturnOverrides.response_code)
}

// string response_error = 2;
inline void ReturnOverrides::clear_response_error() {
  response_error_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& ReturnOverrides::response_error() const {
  // @@protoc_insertion_point(field_get:coprocess.ReturnOverrides.response_error)
  return response_error_.GetNoArena();
}
inline void ReturnOverrides::set_response_error(const ::std::string& value) {
  
  response_error_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:coprocess.ReturnOverrides.response_error)
}
#if LANG_CXX11
inline void ReturnOverrides::set_response_error(::std::string&& value) {
  
  response_error_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:coprocess.ReturnOverrides.response_error)
}
#endif
inline void ReturnOverrides::set_response_error(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  response_error_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:coprocess.ReturnOverrides.response_error)
}
inline void ReturnOverrides::set_response_error(const char* value, size_t size) {
  
  response_error_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:coprocess.ReturnOverrides.response_error)
}
inline ::std::string* ReturnOverrides::mutable_response_error() {
  
  // @@protoc_insertion_point(field_mutable:coprocess.ReturnOverrides.response_error)
  return response_error_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* ReturnOverrides::release_response_error() {
  // @@protoc_insertion_point(field_release:coprocess.ReturnOverrides.response_error)
  
  return response_error_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void ReturnOverrides::set_allocated_response_error(::std::string* response_error) {
  if (response_error != NULL) {
    
  } else {
    
  }
  response_error_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), response_error);
  // @@protoc_insertion_point(field_set_allocated:coprocess.ReturnOverrides.response_error)
}

// map<string, string> headers = 3;
inline int ReturnOverrides::headers_size() const {
  return headers_.size();
}
inline void ReturnOverrides::clear_headers() {
  headers_.Clear();
}
inline const ::google::protobuf::Map< ::std::string, ::std::string >&
ReturnOverrides::headers() const {
  // @@protoc_insertion_point(field_map:coprocess.ReturnOverrides.headers)
  return headers_.GetMap();
}
inline ::google::protobuf::Map< ::std::string, ::std::string >*
ReturnOverrides::mutable_headers() {
  // @@protoc_insertion_point(field_mutable_map:coprocess.ReturnOverrides.headers)
  return headers_.MutableMap();
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace coprocess

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_INCLUDED_coprocess_5freturn_5foverrides_2eproto
