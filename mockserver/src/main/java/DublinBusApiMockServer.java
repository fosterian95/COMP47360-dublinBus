import org.mockserver.configuration.ConfigurationProperties;
import org.mockserver.integration.ClientAndServer;

import static org.mockserver.integration.ClientAndServer.startClientAndServer;

public class DublinBusApiMockServer {

    public DublinBusApiMockServer() {
        // support CORS https://mock-server.com/mock_server/CORS_support.html
        ConfigurationProperties.enableCORSForAllResponses(true);
        ConfigurationProperties.corsAllowOrigin("*");
        // save api settings in a json file https://mock-server.com/mock_server/persisting_expectations.html
        ConfigurationProperties.persistExpectations(true);
        ConfigurationProperties.persistedExpectationsPath("MockServerInit.json");
        ConfigurationProperties.initializationJsonPath("init.json");
        // start a MockServer by Java client API https://mock-server.com/mock_server/running_mock_server.html#client_api
        ClientAndServer clientAndServer = startClientAndServer(1080);
    }

    public static void main(String[] args) {
        new DublinBusApiMockServer();
    }
}
