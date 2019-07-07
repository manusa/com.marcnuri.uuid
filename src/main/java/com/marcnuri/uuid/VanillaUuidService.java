/*
 * VanillaUuidService.java
 *
 * Created on 2019-07-07, 16:46
 */
package com.marcnuri.uuid;

import java.util.Collection;
import java.util.UUID;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import org.springframework.stereotype.Service;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
@Service
public class VanillaUuidService implements UuidService {

  @Override
  public UUID getUuid() {
    return UUID.randomUUID();
  }

  @Override
  public Collection<UUID> getUuid(int quantity) {
    return Stream.generate(this::getUuid).limit(quantity).collect(Collectors.toList());
  }
}
