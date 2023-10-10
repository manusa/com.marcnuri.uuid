/*
 * UuidResourceTest.java
 *
 * Created on 2019-07-07, 17:29
 */
package com.marcnuri.uuid;

import static org.mockito.Mockito.doReturn;
import static org.springframework.http.MediaType.APPLICATION_JSON_UTF8;

import java.util.Arrays;
import java.util.UUID;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.mockito.Mockito;
import org.springframework.test.web.reactive.server.WebTestClient;
import org.springframework.test.web.reactive.server.WebTestClient.ResponseSpec;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
public class UuidResourceTest {

  private UuidService mockUuidService;
  private WebTestClient webTestClient;

  @Before
  public void setUp() {
    mockUuidService = Mockito.mock(UuidService.class);

    webTestClient = WebTestClient
        .bindToController(new UuidResource(mockUuidService))
        .configureClient()
        .build();
  }

  @After
  public void tearDown() {
    webTestClient = null;
    mockUuidService = null;
  }

  @Test
  public void getUuid_shouldReturnOk() {
    // Given
    doReturn(
        new UUID(0L, 0L)
    ).when(mockUuidService).getUuid();
    // When
    final ResponseSpec result = webTestClient.get()
        .uri("/")
        .accept(APPLICATION_JSON_UTF8)
        .exchange();
    // Then
    result.expectStatus().isOk();
    result.expectBody(String.class).isEqualTo("00000000-0000-0000-0000-000000000000");
  }

  @Test
  public void getUuid_invalidNegativeQuantity_shouldReturnBadRequest() {
    // Given
    doReturn(
        new UUID(0L, 0L)
    ).when(mockUuidService).getUuid();
    // When
    final ResponseSpec result = webTestClient.get()
        .uri("/?quantity=-1")
        .accept(APPLICATION_JSON_UTF8)
        .exchange();
    // Then
    result.expectStatus().isBadRequest();
  }

  @Test
  public void getUuid_invalidPositive_shouldReturnBadRequest() {
    // Given
    doReturn(
        new UUID(0L, 0L)
    ).when(mockUuidService).getUuid();
    // When
    final ResponseSpec result = webTestClient.get()
        .uri("/?quantity=20001")
        .accept(APPLICATION_JSON_UTF8)
        .exchange();
    // Then
    result.expectStatus().isBadRequest();
  }

  @Test
  public void getUuid_validQuantity_shouldReturnOk() {
    // Given
    doReturn(
        Arrays.asList(
          new UUID(0L, 0L),
          new UUID(0L, 0L)
        )
    ).when(mockUuidService).getUuid(2);
    // When
    final ResponseSpec result = webTestClient.get()
        .uri("/?quantity=2")
        .accept(APPLICATION_JSON_UTF8)
        .exchange();
    // Then
    result.expectStatus().isOk();
    result.expectBody().jsonPath("$[0]")
        .isEqualTo("00000000-0000-0000-0000-000000000000");
  }
}
