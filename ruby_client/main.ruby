this_dir = File.expand_path(File.dirname(__FILE__))
proto_dir = File.join(this_dir, 'proto')
$LOAD_PATH.unshift(proto_dir) unless $LOAD_PATH.include?(proto_dir)

require 'grpc'
require 'server_services_pb'

def main
  stub = Proto::UserInfo::Stub.new('localhost:8080', :this_channel_is_insecure)
  
  id = ARGV.size > 0 ?  ARGV[0] : "1"
  message = stub.get_user(Proto::UserRequest.new(id: id)).message
  p "Info: #{message}"
end

main
