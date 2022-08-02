//
// Created by Homin Su on 2021/8/16.
//

#include <fstream>

#include "get_cred.h"
#include "logger.h"

std::string GetCred::GetFileContents(const std::string &_path) {
  std::ifstream file_stream(_path);
  if (!file_stream.good()) {
    FATAL("Open Cert File Failed");
  }
  std::string contents;
  contents.assign((std::istreambuf_iterator<char>(file_stream)), std::istreambuf_iterator<char>());
  file_stream.close();
  return contents;
}

std::shared_ptr<grpc::ServerCredentials> GetCred::GetServerCred() {
  auto root_cert = GetFileContents("/cert/ca.crt");
  auto key_str = GetFileContents("/cert/server.key");
  auto cert_str = GetFileContents("/cert/server.pem");
  auto x509KeyPair = grpc::SslServerCredentialsOptions::PemKeyCertPair{key_str, cert_str};

  grpc::SslServerCredentialsOptions cred_option;
  cred_option.pem_root_certs = root_cert;
  cred_option.pem_key_cert_pairs.push_back(x509KeyPair);

  return grpc::SslServerCredentials(cred_option);
}

std::shared_ptr<grpc::ChannelCredentials> GetCred::GetClientCred() {
  auto root_cert = GetFileContents("/cert/ca.crt");
  auto key_str = GetFileContents("/cert/client.key");
  auto cert_str = GetFileContents("/cert/client.pem");

  grpc::SslCredentialsOptions cred_option;
  cred_option.pem_root_certs  = root_cert;
  cred_option.pem_private_key = key_str;
  cred_option.pem_cert_chain  = cert_str;

  return grpc::SslCredentials(cred_option);
}
