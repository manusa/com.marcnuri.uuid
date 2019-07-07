/*
 * UuidConfiguration.java
 *
 * Created on 2019-07-07, 15:57
 */
package com.marcnuri.uuid;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.reactive.config.CorsRegistry;
import org.springframework.web.reactive.config.EnableWebFlux;
import org.springframework.web.reactive.config.WebFluxConfigurer;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
@Configuration
@ComponentScan("com.marcnuri.uuid")
@EnableWebFlux
public class UuidConfiguration {

  @Value("${allowedOrigins:}")
  private String[] allowedOrigins;

  @Bean
  public WebFluxConfigurer corsConfigurer() {
    return new WebFluxConfigurer() {
      @Override
      public void addCorsMappings(CorsRegistry registry) {
        registry.addMapping("/**").allowedOrigins(allowedOrigins);
      }
    };
  }
}
