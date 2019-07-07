/*
 * UuidResource.java
 *
 * Created on 2019-07-07, 16:51
 */
package com.marcnuri.uuid;

import java.util.List;
import java.util.UUID;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
@RestController
@RequestMapping(path = "/")
public class UuidResource {

  private static final int BULK_MIN = 1;
  private static final int BULK_MAX = 20000;

  private final UuidService uuidService;

  public UuidResource(UuidService uuidService) {
    this.uuidService = uuidService;
  }

  @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
  public Mono<String> getUuid() {
    return Mono.just(uuidService.getUuid()).map(UUID::toString);
  }

  @GetMapping(params = {"quantity"}, produces = MediaType.APPLICATION_JSON_VALUE)
  public Mono<List<String>> getUuid(@RequestParam(name="quantity") int quantity) {
    if (quantity < BULK_MIN || quantity > BULK_MAX) {
      throw new ResponseStatusException(HttpStatus.BAD_REQUEST,
          String.format("Quantity must be betweern %s and %s", BULK_MIN, BULK_MAX));
    }
    return Flux.fromIterable(uuidService.getUuid(quantity))
        .map(UUID::toString)
        .collectList();
  }
}
