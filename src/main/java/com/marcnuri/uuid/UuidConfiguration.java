/*
 * UuidConfiguration.java
 *
 * Created on 2019-07-07, 15:57
 */
package com.marcnuri.uuid;

import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.reactive.config.EnableWebFlux;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
@Configuration
@ComponentScan("com.marcnuri.uuid")
@EnableWebFlux
public class UuidConfiguration {

}
