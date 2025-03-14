generate_buf:
	cd ./buf_schema && buf generate
	echo $?
setup_buf:
	cd ./buf_schema && cp -r ./gen/* ../feed_collector/feed_collector/gen
	cd ./buf_schema && cp -r ./gen/* ../federation_orchestrator/federation_orchestrator/gen



