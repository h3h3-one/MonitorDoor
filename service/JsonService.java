package service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import models.ConfigurationModelDoorAndMonitor;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;

public class JsonService {
    private static final Logger LOG = LogManager.getLogger(JsonService.class);

    public void isNewFile(File file) {
        try {
            if (file.createNewFile()) {
                FileOutputStream out = new FileOutputStream(file);

                ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
                String json = ow.writeValueAsString(new ConfigurationModelDoorAndMonitor());

                out.write(json.getBytes());
                out.close();

                LOG.info("Файл конфигурации успешно создан. Запустите программу заново.  ПУТЬ: " + file.getAbsolutePath());
                System.exit(0);
            }
        } catch (IOException e) {
            LOG.error("Ошибка: " + e.getMessage());
        }
    }

    public ConfigurationModelDoorAndMonitor getConfigParam() {
        ConfigurationModelDoorAndMonitor model = null;
        try {
            model = new ConfigurationModelDoorAndMonitor();
            ObjectMapper mapper = new ObjectMapper();
            File file = new File("MonitorDoorConfig.json");
            isNewFile(file);
            model = mapper.readValue(file, ConfigurationModelDoorAndMonitor.class);
        } catch (IOException e) {
            LOG.error("Ошибка: " + e.getMessage());
        }
        return model;
    }
}
