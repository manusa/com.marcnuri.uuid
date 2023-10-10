package com.marcnuri.uuid;

import java.util.Collection;
import java.util.UUID;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
public interface UuidService {

  UUID getUuid();

  Collection<UUID> getUuid(int quantity);
}
